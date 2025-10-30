package api

import (
	"blog/consts"
	"blog/global"
	"blog/models"
	"blog/res"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

type UserDataResponse struct {
	Fans   int `json:"fans"`   // 粉丝数量
	Follow int `json:"follow"` // 关注数
	Likes  int `json:"likes"`  // 获赞数

}

func NewUserDataResponse(fans int, follow int, likes int) *UserDataResponse {
	return &UserDataResponse{Fans: fans, Follow: follow, Likes: likes}
}

// GetUserDataView TODO 修改返回值, 上面这个从登录的时候获取
func (UserApi) GetUserDataView(c *gin.Context) {
	userId, _ := c.Get(consts.UserId)
	db := global.MysqlDB
	var user models.User
	// 获取用户
	db.Where("id = ?", userId).Take(&user)
	if user.Username == "" {
		res.Fail(c, 500, consts.UserNotFound)
	}
	// 封装用户信息
	userDataResponse := NewUserDataResponse(12, 103, 1022)
	res.Success(c, userDataResponse, "")
}
