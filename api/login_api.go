package api

import (
	"blog/consts"
	"blog/enum"
	"blog/global"
	"blog/models"
	"blog/res"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginApi struct {
}

type LoginRequest struct {
	LoginType enum.LoginType `json:"loginType"`

	// 账号密码登录
	// omitempty序列化时如果为空忽略该字段
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`

	// 邮箱登录
	Email     string `json:"email,omitempty"`
	EmailCode string `json:"emailCode,omitempty"`
}

type LoginResponse struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

func NewLoginResponse(userID uint, username string, nickname string, avatar string, token string) *LoginResponse {
	return &LoginResponse{UserID: userID, Username: username, Nickname: nickname, Avatar: avatar, Token: token}
}

func (LoginApi) LoginView(c *gin.Context) {
	//c.ShouldBindJSON()
	var loginRequest LoginRequest
	// 解析请求体json参数
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	db := global.MysqlDB
	// 保存用户信息
	var user models.User
	switch loginRequest.LoginType {
	case enum.PasswordLoginType:
		if loginRequest.Username == "" || loginRequest.Password == "" {
			res.Fail(c, http.StatusBadRequest, consts.UsernameOrPwdNull)
			return
		}
		// 根据用户名查询该用户
		db.Take(&user, "username = ?", loginRequest.Username)
		// 校验用户是否存在
		if user.ID == 0 {
			res.Fail(c, http.StatusBadRequest, consts.UserNotFound)
			return
		}
		// 校验用户名密码
		// 获取加密密码
		hash, _ := bcrypt.GenerateFromPassword([]byte(loginRequest.Password), bcrypt.DefaultCost)
		// 判断密码是否正确
		err := bcrypt.CompareHashAndPassword(hash, []byte(loginRequest.Password))
		if err != nil {
			res.Fail(c, http.StatusBadRequest, consts.PwdError)
			return
		}
	case enum.EmailLoginType:
		if loginRequest.Email == "" || loginRequest.EmailCode == "" {
			res.Fail(c, http.StatusBadRequest, consts.EmailOrCodeNull)
			return
		}
		// 判断邮箱是否存在
		db.Take(&user, "email = ?", loginRequest.Email)
		if user.ID == 0 {
			res.Fail(c, http.StatusBadRequest, consts.UserNotFound)
			return
		}
		// 校验验证码是否正确
		code, err := global.RedisDB.Get(global.Ctx, consts.EmailCodeKeyPrefix+user.Email).Result()
		if err != nil || code != loginRequest.EmailCode {
			res.Fail(c, http.StatusBadRequest, consts.CodeError)
			return
		}
	default:
		res.Fail(c, http.StatusBadRequest, consts.UnknowLoginType)
		return
	}
	// 生成token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		res.Fail(c, http.StatusInternalServerError, consts.GenJwtError)
	}
	// 生成用户信息结构体返回
	var loginResponse = NewLoginResponse(user.ID, user.Username, user.Nickname, user.Avatar, token)
	res.Success(c, loginResponse, consts.LoginSuccess)
}
