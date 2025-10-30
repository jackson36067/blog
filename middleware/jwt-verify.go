package middleware

import (
	"blog/consts"
	"blog/res"
	"blog/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func JwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Authorization 头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			res.Fail(c, 401, consts.TokenNull)
			c.Abort()
			return
		}
		// 处理 Bearer 前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
			authHeader = parts[1]
		}
		// 解析token
		claims, err := utils.ParseToken(authHeader)
		if err != nil {
			res.Fail(c, 401, consts.InvalidToken)
			c.Abort()
			return
		}
		// 检查是否过期
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
			res.Fail(c, 401, consts.ExpireToken)
			c.Abort()
			return
		}
		logrus.Infof("token校验成功: %s", authHeader)
		c.Set(consts.UserId, claims.UserID)
		c.Next()
	}
}
