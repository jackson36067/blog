package api

import (
	"blog/consts"
	"blog/global"
	"blog/models"
	"blog/res"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterApi struct{}

type RegisterRequest struct {
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	EmailCode string `json:"emailCode"`
}

func (RegisterApi) RegisterView(c *gin.Context) {
	// 1. 解析请求体json参数
	var registerRequest RegisterRequest
	db := global.MysqlDB
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		res.Fail(c, http.StatusBadRequest, err.Error())
	}
	// 2. 判断用户名以及邮箱是否已经被使用
	var user models.User
	db.Where("username = ?", registerRequest.Username).Take(&user)
	// 用户名已经被使用
	if user.Username != "" {
		res.Fail(c, http.StatusBadRequest, consts.UsernameExist)
		return
	}
	// 邮箱已经使用
	db.Where("email = ?", registerRequest.Email).Take(&user)
	if user.Username != "" {
		res.Fail(c, http.StatusBadRequest, consts.EmailExist)
		return
	}
	// 3. 校验验证码是否正确
	val, err := global.RedisDB.Get(global.Ctx, consts.EmailCodeKeyPrefix+registerRequest.Email).Result()
	if err != nil || val != registerRequest.EmailCode {
		res.Fail(c, http.StatusBadRequest, consts.CodeError)
		return
	}
	// 生成新用户
	user.Username = registerRequest.Username
	user.Nickname = registerRequest.Nickname
	// 生成加密密码
	hash, _ := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	user.Email = registerRequest.Email
	db.Create(&user)
	res.Success(c, nil, consts.RegisterSuccess)
}
