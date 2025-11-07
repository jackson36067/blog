package routers

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

func CategoryRouter(router *gin.RouterGroup) {
	app := api.App.CategoryApi
	// 不需要登录就可以访问的路由
	public := router.Group("/category")
	public.GET("/list", app.GetCategoryList)
	// 需要登录才可以访问的路由
	//private := router.Group("/category")
}
