package main

import (
	"context"
	"editortest/pkg/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("v1")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithConfig(logger.LoggerConfig))
	// r := gin.Default()
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.GET("/ping", func(c *gin.Context) {
		key := logger.ContextKey("username")
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), key, "fabio kälin asdf"))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/ping/abc", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongabc",
		})
	})
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "1",
		})
	})
	r.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "1",
		})
	})
	r.HandleMethodNotAllowed = true

	r.Run(":8000")
}
