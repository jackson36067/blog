package api

import (
	"blog/consts"
	"blog/core"
	"blog/dto/response"
	"blog/global"
	"blog/models"
	"blog/res"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// GetUserDataView 获取用户信息
func (UserApi) GetUserDataView(c *gin.Context) {
	userId, _ := c.Get(consts.UserId)
	db := global.MysqlDB
	var user models.User
	// 获取用户
	db.Preload("Articles").Where("id = ?", userId).Take(&user)
	if user.Username == "" {
		res.Fail(c, 500, consts.UserNotFound)
	}
	// 封装用户信息
	// 获取用户粉丝以及关注数量
	var followers []uint
	var followed []uint
	db.Model(&models.UserFollow{}).Where("follower_id = ?", userId).Pluck("followed_id", &followed)
	db.Model(&models.UserFollow{}).Where("followed_id = ?", userId).Pluck("follower_id", &followers)
	// 获取ip地址
	core.InitIPDB()
	ip, err := core.GetIpAddress(c.ClientIP())
	if err != nil {
		res.Fail(c, 500, consts.IpParseError)
	}
	userDataResponse := response.UserDataResponse{
		OriginArticle: len(user.Articles),
		Fans:          len(followers),
		Follow:        len(followed),
		IP:            ip,
		JoinTime:      user.CreatedAt.Format("2006-01-02 15:04:05"),
		CodeAge:       user.CodeAge,
		Avatar:        user.Avatar,
	}
	res.Success(c, userDataResponse, "")
}
