package api

import (
	"math"
	"sort"

	"blog/consts"
	"blog/dto/request"
	"blog/dto/response"
	"blog/enum"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

type MessageApi struct{}

// GetMessageHistory 获取用户聊天记录
func (MessageApi) GetMessageHistory(c *gin.Context) {
	userIdAny, _ := c.Get(consts.UserId)
	// 用户id
	userId := userIdAny.(uint)
	// 聊天用户id
	chatUserIdStr := c.Param("id")
	chatUserId, _ := utils.StringToUint(chatUserIdStr)
	// 获取当前用户与聊天用户的聊天记录
	db := global.MysqlDB
	var messages []models.Message
	db.Preload("SendUser").
		Preload("ReceiveUser").
		Where("status = ? AND type = ?", enum.Normal, enum.PrivateMessage).
		Where(db.Where("send_user_id = ? AND receive_user_id = ?", userId, chatUserId).
			Or("send_user_id = ? AND receive_user_id = ?", chatUserId, userId)).
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
}

// GetOtherUserMessage 获取其他用户的消息
func (MessageApi) GetOtherUserMessage(c *gin.Context) {
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
