package routers

import (
	"blog/api"
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func SessionRouter(router *gin.RouterGroup) {
	app := api.App.SessionApi
	private := router.Group("/session")
	private.Use(middleware.JwtVerify())
	private.GET("/history", app.GetUserSessionHistory)
	private.PUT("/update/:id", app.UpdateSession)
}
