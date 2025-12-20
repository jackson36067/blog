package service

import (
	"sort"

	"blog/dto/response"
	"blog/enum"
	"blog/global"
	"blog/models"
	"blog/utils"
)

func BuildChatListItems(userID uint, sessions []models.Session) ([]response.SessionResponse, error) {
	var list []response.SessionResponse

	for _, s := range sessions {
		var chatUserID uint
		var isPinned bool
		var isMuted bool
		var chatUser models.User

		// 判断当前用户是 A 还是 B
		if s.UserIDA == userID {
			chatUserID = s.UserIDB
			isPinned = s.IsPinnedA
			isMuted = s.IsMutedA
			chatUser = s.UserB
		} else {
			chatUserID = s.UserIDA
			isPinned = s.IsPinnedB
			isMuted = s.IsMutedB
			chatUser = s.UserA
		}

		// 判断用户是否关注聊天用户
		db := global.MysqlDB
		var followed models.UserFollow
		db.Where("followed_id = ? AND follower_id = ?", chatUserID, userID).First(&followed)

		// 获取当前聊天队列未读消息数目
		var unReadMessageCount int64
		db.Model(&models.Message{}).
			Where("send_user_id = ? AND receive_user_id = ? AND type = ?", chatUserID, userID, enum.PrivateMessage).
			Count(&unReadMessageCount)

		item := response.SessionResponse{
			SessionID:      s.ID,
			ChatUserID:     chatUserID,
			ChatUsername:   chatUser.Username,
			ChatUserAvatar: chatUser.Avatar,
			LatestMessage:  s.LatestMessage,
			LatestChatTime: utils.FormatChatTime(s.LatestChatTime),
			IsPinned:       isPinned,
			IsMuted:        isMuted,
			IsFollow:       followed.ID > 0,
			UnreadCount:    uint(unReadMessageCount),
		}

		list = append(list, item)
	}

	// 在这里排序（置顶优先, 且不改变一开始的顺序）
	sort.Slice(list, func(i, j int) bool {
		return list[i].IsPinned && !list[j].IsPinned
	})

	return list, nil
}
