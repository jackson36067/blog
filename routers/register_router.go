package routers

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup) {
	app := api.App.RegisterApi
	router.POST("/register", app.RegisterView)
}
