package routers

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

func LoginRouter(router *gin.RouterGroup) {
	app := api.App.LoginApi
	router.POST("/login", app.LoginView)
}
