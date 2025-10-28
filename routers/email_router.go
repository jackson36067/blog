package routers

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

func EmailRouter(router *gin.RouterGroup) {
	app := api.App.EmailApi
	router.GET("/email/code", app.SendEmailCodeView)
}
