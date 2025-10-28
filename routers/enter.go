package routers

import (
	"blog/global"

	"github.com/gin-gonic/gin"
)

func Run() {
	// 初始化路由
	router := gin.Default()
	// 挂载路由并进行路由分组
	// 1.登录, 注册, 发送验证码路由,
	loginOrRegisterRouterGroup := router.Group("/api")
	// 登录路由
	LoginRouter(loginOrRegisterRouterGroup)
	// 注册路由
	RegisterRouter(loginOrRegisterRouterGroup)
	// 发送验证码路由
	EmailRouter(loginOrRegisterRouterGroup)
	// 2.不需要登录就可以访问的路由
	// 3.需要登录访问的路由
	systemConf := global.Conf.System
	// 启动路由
	router.Run(systemConf.Host + ":" + systemConf.Port)
}
