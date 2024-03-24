package configs

import (
	"github.com/gin-gonic/gin"
	"githut.com/goredan/server/controllers"
)

func SetupRouter(router *gin.Engine) {
	api := router.Group("/api")
	ws := router.Group("/ws")

	homeController := controllers.HomeController{}

	api.GET("", homeController.Home)
	api.GET("/ping", homeController.Ping)

	wsController := controllers.WebSocketController{}
	ws.GET("", wsController.HandleConnections)
}
