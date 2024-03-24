package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"githut.com/goredan/server/configs"
)

func main() {
	// Load config
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Router
	gin.SetMode(config.GinMode)
	router := gin.Default()

	// CORS
	corsConf := cors.DefaultConfig()
	corsConf.AllowOrigins = []string{config.ClientSite}
	router.Use(cors.New(corsConf))

	configs.SetupRouter(router)

	// Start Server
	fmt.Println("Server running on ", config.ServerHost+":"+config.ServerPort)
	if err := router.Run(config.ServerHost + ":" + config.ServerPort); err != nil {
		panic(err)
	}
}
