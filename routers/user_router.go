package routers

import (
	"blog/api"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	app := api.App.UserApi
	// 用户路由不需要登录的路由组
	public := router.Group("/user")
	public.POST("/upload", app.Upload)
	// 用户路由需要登录的路由组
	private := router.Group("/user")
	private.Use(middleware.JwtVerify())
	private.GET("/data", app.GetUserDataView)
	private.GET("/achievement", app.GetUserAchievementListView)
	private.GET("/likes", app.GetUserLikeArticlesView)
	private.GET("/browse/history", app.GetUserBrowseArticleHistoryView)
	private.GET("/followed", app.GetUserFollowedView)
	private.POST("/follow/:id", app.UpdateFollowView)
	private.GET("/follower", app.GetUserFollowerView)
	private.GET("/comment", app.GetUserCommentsView)
	private.PUT("/update", app.UpdateUserInfoView)
	private.GET("/login/log", app.GetUserLoginLogPaginationView)
}
