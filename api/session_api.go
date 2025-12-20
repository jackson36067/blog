package api

import (
	"blog/consts"
	"blog/dto/request"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/service"
	"blog/utils"

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

// UpdateSession 更新会话信息
func (SessionApi) UpdateSession(c *gin.Context) {
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
