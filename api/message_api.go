package api

import (
	"log"
	"math"
	"sort"
	"time"

	"blog/consts"
	"blog/dto/request"
	"blog/dto/response"
	"blog/enum"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/utils"
	"blog/ws"

	"github.com/gin-gonic/gin"
)

type MessageApi struct{}

// GetMessageHistoryView 获取用户聊天记录
func (MessageApi) GetMessageHistoryView(c *gin.Context) {
	// 聊天用户id
	sessionId := c.Param("id")
	// 获取当前用户与聊天用户的聊天记录
	db := global.MysqlDB
	var messages []models.Message
	db.Preload("SendUser").
		Preload("ReceiveUser").
		Where("status = ? AND type = ? AND session_id = ?", enum.Normal, enum.PrivateMessage, sessionId).
		Order("send_time ASC").
		Find(&messages)
	// 按发送时间（精确到分钟）分组
	groupedMessages := make(map[string][]response.ChatResponse)
	for _, message := range messages {
		// 将时间格式化为"YYYY-MM-DD HH:MM"格式，精确到分钟
		minuteKey := message.SendTime.Format("2006-01-02 15:04")
		// 判断是否已经存在该分组
		if _, exists := groupedMessages[minuteKey]; !exists {
			groupedMessages[minuteKey] = []response.ChatResponse{}
		}
		// 添加消息到分组
		chatResp := response.ChatResponse{
			ID:             message.ID,
			UserID:         message.SendUserID,
			Username:       message.SendUser.Username,
			UserAvatar:     message.SendUser.Avatar,
			ChatUserID:     message.ReceiveUserID,
			ChatUsername:   message.ReceiveUser.Username,
			ChatUserAvatar: message.ReceiveUser.Avatar,
			Message:        message.Content,
			SendTime:       message.SendTime.Format("2006-01-02 15:04:05"),
		}
		if _, exists := groupedMessages[minuteKey]; !exists {
			groupedMessages[minuteKey] = []response.ChatResponse{}
		}
		groupedMessages[minuteKey] = append(groupedMessages[minuteKey], chatResp)
	}
	// 将分组后的数据转换为前端需要的格式
	var chatMessageHistory []response.GroupedChatResponse
	// 遍历分组并按时间顺序排列
	var minuteKeys []string
	for key := range groupedMessages {
		minuteKeys = append(minuteKeys, key)
	}
	sort.Strings(minuteKeys)
	// 按时间排序（从早到晚）放入最终的分组数组
	for _, minuteKey := range minuteKeys {
		group := response.GroupedChatResponse{
			SendTime: minuteKey,
			Messages: groupedMessages[minuteKey],
		}
		chatMessageHistory = append(chatMessageHistory, group)
	}
	res.Success(c, chatMessageHistory, "")
	userIdAny, _ := c.Get(consts.UserId)
	userId, _ := userIdAny.(uint)
	// 清空该用户的对该会话的未读消息
	go func(sessionId string, userId uint) {
		db.Model(&models.Message{}).Where("session_id = ? AND receive_user_id = ?", sessionId, userId).Update("is_read", true)
	}(sessionId, userId)
}

// GetOtherUserMessageView 获取用户其他的消息
func (MessageApi) GetOtherUserMessageView(c *gin.Context) {
	userIdAny, _ := c.Get(consts.UserId)
	userId, _ := userIdAny.(uint)
	// 解析请求参数
	var params request.OtherMessageRequestParams
	if err := c.ShouldBindQuery(&params); err != nil {
		res.Fail(c, 400, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	// 根据类型获取
	var total int64
	db.Model(&models.Message{}).Where("status = ? AND receive_user_id = ? AND type = ?", enum.Normal, userId, params.Type).Count(&total)

	// 分页查询
	var messages []models.Message
	db.Preload("SendUser").
		Where("status = ? AND receive_user_id = ? AND type = ?", enum.Normal, userId, params.Type).
		Order("send_time ASC").
		Offset((params.Page - 1) * params.PageSize).
		Limit(params.PageSize).
		Find(&messages)
	otherMessagePagination := utils.MapSlice(messages, func(message models.Message) response.OtherMessageResponse {
		var isFollow bool
		var title string
		if params.Type == enum.FanMessage {
			var follow models.UserFollow
			db.Where("followed_id = ? AND follower_id = ?", message.SendUserID, userId).First(&follow)
			isFollow = follow.ID > 0
		} else {
			// 获取评论/点赞/收藏消息的标题
			db.Model(&models.Article{}).Where("id = ?", message.Extra).Pluck("title", &title)
		}
		return response.OtherMessageResponse{
			ID:            message.ID,
			UserID:        message.SendUserID,
			Username:      message.SendUser.Username,
			UserAvatar:    message.SendUser.Avatar,
			Message:       message.Content,
			ActionMessage: message.ActionMessage,
			SendTime:      utils.FormatMessageTime(message.SendTime),
			Extra:         message.Extra,
			Title:         title,
			IsFollow:      isFollow,
		}
	})
	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))
	pagination := res.Pagination{
		Page:          params.Page,
		PageSize:      params.PageSize,
		TotalElements: total,
		TotalPages:    totalPages,
		Data:          otherMessagePagination,
	}
	res.Success(c, pagination, "")
}

// DeleteMessageView 删除消息
func (MessageApi) DeleteMessageView(c *gin.Context) {
	var params request.RemoveMessageRequestParams
	userId, _ := c.Get(consts.UserId)
	if err := c.ShouldBindJSON(&params); err != nil {
		res.Fail(c, 400, consts.RequestParamParseError)
	}
	db := global.MysqlDB
	// 删除消息
	err := db.Delete(&models.Message{}, "id in ? AND receive_user_id = ?", params.MessageIdList, userId).Error
	if err != nil {
		res.Fail(c, 400, consts.DeleteMessageError)
	}
	res.Success(c, nil, consts.DeleteMessageSuccess)
}

// SendMessageView 发送消息
func (MessageApi) SendMessageView(c *gin.Context) {
	var params request.SendMessageRequestParams
	userIdAny, _ := c.Get(consts.UserId)
	userId, _ := userIdAny.(uint)
	if err := c.ShouldBindJSON(&params); err != nil {
		res.Fail(c, 400, consts.RequestParamParseError)
	}
	tx := global.MysqlDB.Begin()
	// 新增消息
	message := &models.Message{
		Type:          enum.PrivateMessage,
		SendUserID:    userId,
		ReceiveUserID: params.UserID,
		Content:       params.Content,
		ContentType:   params.ContentType,
		Status:        enum.Normal,
		IsRead:        false,
		SendTime:      time.Now(),
		SessionID:     &params.SessionID,
	}
	if err := tx.Create(message).Error; err != nil {
		res.Fail(c, 400, err.Error())
		tx.Rollback()
		return
	}
	// 更改会话状态
	if err := tx.Model(&models.Session{}).
		Where("id = ?", params.SessionID).
		Updates((map[string]any{
			"latest_message":   message.Content,
			"latest_chat_time": message.SendTime,
			"deleted_at_a":     nil,
			"deleted_at_b":     nil,
		})).Error; err != nil {
		res.Fail(c, 400, err.Error())
		tx.Rollback()
		return
	}

	wsm := ws.WsManager
	userClient := wsm.GetClient(userId)
	chatUserClient := wsm.GetClient(params.UserID)

	// 至少有一方在线
	if userClient == nil && chatUserClient == nil {
		return
	}

	// 查用户信息（只查一次）
	var user, chatUser models.User
	tx.First(&user, "id = ?", userId)
	tx.First(&chatUser, "id = ?", params.UserID)

	// 构建 WS 消息
	wsMessage := response.PrivateSendMessageResponse{
		SendTime: utils.FormatChatTime(message.SendTime),
		Message: response.ChatResponse{
			ID:             message.ID,
			UserID:         userId,
			Username:       user.Username,
			UserAvatar:     user.Avatar,
			ChatUserID:     userId,
			ChatUsername:   chatUser.Username,
			ChatUserAvatar: chatUser.Avatar,
			Message:        params.Content,
			SendTime:       message.SendTime.Format("2006-01-02 15:04:05"),
		},
		SessionID: params.SessionID,
		ChatTime:  utils.FormatChatTime(message.SendTime),
	}

	// 分别发送
	if userClient != nil {
		_ = wsm.SendToUser(userId, response.SendMessageResponse{
			Type: "chat",
			Data: wsMessage,
		})
		log.Printf("发送给用户%d消息成功: %s", userClient.UserId, params.Content)
	}
	if chatUserClient != nil {
		_ = wsm.SendToUser(params.UserID, response.SendMessageResponse{
			Type: "chat",
			Data: wsMessage,
		})
		log.Printf("发送给用户%d消息成功: %s", chatUserClient.UserId, params.Content)
	}
	tx.Commit()
}
