package routers

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

func ArticleTagRouter(router *gin.RouterGroup) {
	app := api.App.ArticleTagApi
	// 不需要登录的路由
	public := router.Group("/tag")
	public.GET("/hot", app.GetHotArticleTagPagination)
}
