package api

import (
	"blog/consts"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/service"

	"github.com/gin-gonic/gin"
)

type SessionApi struct{}

func (SessionApi) GetUserSessionHistory(c *gin.Context) {
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
		Where("user_id_a = ? OR user_id_b = ?", userId, userId).
		Preload("UserA").
		Preload("UserB").
		Order("latest_chat_time DESC").
		Find(&sessions)
	chatList, _ := service.BuildChatListItems(userId, sessions)
	res.Success(c, chatList, "")
}
