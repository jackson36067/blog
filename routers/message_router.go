package routers

import (
	"blog/api"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func MessageRouter(router *gin.RouterGroup) {
	app := api.App.MessageApi
	private := router.Group("/message")
	private.Use(middleware.JwtVerify())
	private.GET("/chat/:id", app.GetMessageHistoryView)
	private.GET("/other", app.GetOtherUserMessageView)
	private.DELETE("/remove", app.DeleteMessageView)
	private.POST("/send", app.SendMessageView)
}
