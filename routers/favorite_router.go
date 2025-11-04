package routers

import (
	"blog/api"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func FavoriteRouter(router *gin.RouterGroup) {
	app := api.App.FavoriteApi
	// 收藏夹路由不需要登录的路由
	public := router.Group("favorite")
	public.GET("/articles/:id", app.GetFavoriteArticleListView)
	// 收藏夹需要登录的路由
	private := router.Group("/favorite")
	private.Use(middleware.JwtVerify())
	private.GET("/list", app.GetUserFavoriteListView)
	private.POST("/new", app.NewFavoriteView)
	private.PUT("/update/:id", app.UpdateFavorite)
	private.PUT("/move", app.MoveFavoriteArticleView)
	private.DELETE("remove/:id", app.RemoveFavoriteArticle)
	private.DELETE("/delete/:id", app.DeleteFavorite)
}
