package routers

import (
	"blog/api"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	app := api.App.UserApi
	router.GET("/user/data", app.GetUserDataView)
}
