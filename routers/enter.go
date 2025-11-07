package routers

import (
	"blog/global"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func Run() {
	// 初始化路由
	router := gin.Default()
	// 挂载跨域中间件, 对所有路由生效
	router.Use(middleware.Cors())
	// 挂载路由并进行路由分组
	routerGroups := router.Group("/api")
	// 登录路由
	LoginRouter(routerGroups)
	// 注册路由
	RegisterRouter(routerGroups)
	// 发送验证码路由
	EmailRouter(routerGroups)
	// 文章路由
	ArticleRouter(routerGroups)
	// 用户路由
	UserRouter(routerGroups)
	// 收藏级路由
	FavoriteRouter(routerGroups)
	// 文章分类路遇
	CategoryRouter(routerGroups)
	systemConf := global.Conf.System
	// 启动路由
	router.Run(systemConf.Host + ":" + systemConf.Port)
}
