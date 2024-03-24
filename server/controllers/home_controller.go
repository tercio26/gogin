package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct{}

func (home *HomeController) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func (home *HomeController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
