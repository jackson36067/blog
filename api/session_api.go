package api

import (
	"errors"
	"time"

	"blog/consts"
	"blog/dto/request"
	"blog/dto/response"
	"blog/enum"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/service"
	"blog/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SessionApi struct{}

// GetUserSessionHistoryView 获取用户会话历史
func (SessionApi) GetUserSessionHistoryView(c *gin.Context) {
	// 获取用户id
	userAny, _ := c.Get(consts.UserId)
	userId, _ := userAny.(uint)
	if userId == 0 {
		res.Fail(c, 401, "请先登录")
		return
	}
	db := global.MysqlDB
	var sessions []models.Session
	db.Model(&models.Session{}).
		Where(db.Where("user_id_a = ? AND deleted_at_a IS NULL", userId).Or("user_id_b = ? AND deleted_at_b IS NULL", userId)).
		Preload("UserA").
		Preload("UserB").
		Order("latest_chat_time DESC").
		Find(&sessions)
	chatList, _ := service.BuildChatListItems(userId, sessions)
	res.Success(c, chatList, "")
}

// UpdateSessionView 更新会话信息
func (SessionApi) UpdateSessionView(c *gin.Context) {
	// 获取会话id
	sessionIdStr := c.Param("id")
	sessionId, _ := utils.StringToUint(sessionIdStr)
	// 解析请求参数
	var params request.UpdateSessionRequestParams
	if err := c.ShouldBind(&params); err != nil {
		res.Fail(c, 400, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	var session models.Session
	// 根据id查询会话
	db.Where("id = ?", sessionId).First(&session)
	// 获取当前用户id
	userIdAny, ok := c.Get(consts.UserId)
	if !ok {
		res.Fail(c, 401, "请先登录")
		return
	}
	userId, _ := userIdAny.(uint)
	// 封装更新条件
	updates := map[string]any{}
	if params.IsMuted != nil {
		if session.UserIDA == userId {
			updates["is_muted_a"] = *params.IsMuted
		} else if session.UserIDB == userId {
			updates["is_muted_b"] = *params.IsMuted
		}
	}
	if params.IsPinned != nil {
		if session.UserIDA == userId {
			updates["is_pinned_a"] = *params.IsPinned
		} else if session.UserIDB == userId {
			updates["is_pinned_b"] = *params.IsPinned
		}
	}
	if len(updates) > 0 {
		db.Debug().
			Model(&models.Session{}).
			Where("id = ?", sessionId).
			Updates(updates)
	}
}

// DeleteSessionView 删除会话
func (SessionApi) DeleteSessionView(c *gin.Context) {
	sessionId := c.Param("id")
	db := global.MysqlDB

	userId, _ := c.Get(consts.UserId)
	now := time.Now()

	err := db.Transaction(func(tx *gorm.DB) error {
		var session models.Session
		if err := tx.First(&session, "id = ?", sessionId).Error; err != nil {
			return err
		}

		// 标记当前用户删除时间
		if session.UserIDA == userId {
			if err := tx.Model(&session).
				Update("deleted_at_a", now).Error; err != nil {
				return err
			}
			session.DeletedAtA = &now
		} else if session.UserIDB == userId {
			if err := tx.Model(&session).
				Update("deleted_at_b", now).Error; err != nil {
				return err
			}
			session.DeletedAtB = &now
		} else {
			return errors.New("无权操作该会话")
		}

		// 如果双方都已经删除，则清理消息
		if session.DeletedAtA != nil && session.DeletedAtB != nil {
			clearTime := *session.DeletedAtA
			if session.DeletedAtB.Before(clearTime) {
				clearTime = *session.DeletedAtB
			}

			if err := tx.Where(
				"session_id = ? AND send_time < ?",
				session.ID,
				clearTime,
			).Delete(&models.Message{}).Error; err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		res.Fail(c, 500, err.Error())
		return
	}

	res.Success(c, nil, "")
}

// CreateOrGetSession 根据用户名创建或返回会话
func (SessionApi) CreateOrGetSession(c *gin.Context) {
	username := c.Query("username")
	userId, _ := c.Get(consts.UserId)
	db := global.MysqlDB
	// 根据用户名获取用户信息
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		res.Fail(c, 400, consts.UserNotFound)
		return
	}
	// 根据用户id获取会话
	var session models.Session
	now := time.Now()
	if err := db.Where(db.Where("user_id_a = ? AND user_id_b = ?", userId, user.ID).Or("user_id_a = ? AND user_id_b = ?", user.ID, userId)).First(&session).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// 没有这个会话就创建一个(注意属性值需要对应)
		session = models.Session{
			UserIDA:        userId.(uint),
			UserIDB:        user.ID,
			IsPinnedA:      false,
			IsPinnedB:      false,
			IsMutedA:       true, // 默认创建的为免打扰
			IsMutedB:       false,
			LatestChatTime: time.Now(),
			DeletedAtB:     &now, // 默认对方看不到对话栏
		}
		db.Create(&session)
	}
	// 获取当前用户未读消息对于该会话
	var unreadMessageCount int64
	db.Where("receive_id = ? AND session_id = ? AND is_read = ? AND status = ? AND type= ? ", userId, session.ID, false, enum.Normal, enum.PrivateMessage).Count(&unreadMessageCount)
	// 判断当前用户是否关注聊天用户
	var followed models.UserFollow
	db.Where("followed_id = ? AND follower_id = ?", user.ID, userId).First(&followed)
	// 返回创建的会话信息
	sessionResponse := response.SessionResponse{
		SessionID:      session.ID,
		ChatUserID:     user.ID,
		ChatUsername:   user.Username,
		ChatUserAvatar: user.Avatar,
		IsPinned:       false,
		IsMuted:        true, // 默认的为免打扰
		LatestMessage:  "",
		LatestChatTime: utils.FormatMessageTime(now),
		IsFollow:       followed.ID > 0,
		UnreadCount:    uint(unreadMessageCount),
	}
	res.Success(c, sessionResponse, "")
}
