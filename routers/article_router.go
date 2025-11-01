package routers

import (
	"blog/api"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func ArticleRouter(router *gin.RouterGroup) {
	app := api.App.ArticleApi
	// 用户路由不需要登录的路由组
	public := router.Group("/article")
	public.GET("/info", app.GetHomeArticleView)
	// 用户路由需要登录的路由组
	private := router.Group("/article")
	private.Use(middleware.JwtVerify())
	private.GET("/top", app.GetUserTopArticleListView)
}
