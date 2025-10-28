package api

import (
	"blog/consts"
	"blog/global"
	"blog/res"
	"blog/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type EmailApi struct {
}

type SendEmailRequest struct {
	Email string `form:"email"`
}

func (EmailApi) SendEmailCodeView(c *gin.Context) {
	// 解析参数信息
	var sendEmailRequest SendEmailRequest
	if err := c.ShouldBindQuery(&sendEmailRequest); err != nil {
		res.Fail(c, http.StatusBadRequest, err.Error())
	}
	// 发送验证码
	code := utils.GenerateSecureCode()
	err := utils.SendEmail([]string{sendEmailRequest.Email}, "jackson-blog登录验证码", code)
	if err != nil {
		logrus.Errorf("%s %s", consts.SendMailError, err.Error())
		return
	}
	// 将code存入redis, 设置3分钟过期时间
	global.RedisDB.Set(global.Ctx, consts.EmailCodeKeyPrefix+sendEmailRequest.Email, code, 3*time.Minute)
	res.Success(c, nil, "获取验证码成功")
}
