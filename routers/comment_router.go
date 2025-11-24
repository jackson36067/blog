package routers

import (
	"blog/api"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRouter(router *gin.RouterGroup) {
	app := api.App.CommentApi
	// 不需要登录就可以访问的接口
	//public := router.Group("/comment")
	// 需要登录才可以访问的接口
	private := router.Group("/comment")
	private.Use(middleware.JwtVerify())
	private.POST("/publish", app.PublishCommentView)
}
