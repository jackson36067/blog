package routers

import (
	"blog/api"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func ArticleRouter(router *gin.RouterGroup) {
	app := api.App.ArticleApi
	// 文章路由不需要登录的路由组
	public := router.Group("/article")
	public.GET("/info", app.GetHomeArticleView)
	// 获取文章的分类以及标签
	public.GET("/meta", app.GetArticleHotTagsAndRandCategoryView)
	public.GET("/category/list", app.GetArticleCategoryList)
	public.GET("/tag/list", app.GetArticleTagList)
	// 文章路由需要登录的路由组
	private := router.Group("/article")
	private.Use(middleware.JwtVerify())
	private.GET("/top", app.GetUserTopArticleListView)
	private.GET("/my", app.GetUserArticlePaginationView)
	private.GET("/statistic", app.GetUserArticleCreateProcess)
	private.DELETE("/remove/browse", app.ClearUserBrowseArticleHistoryView)
}
