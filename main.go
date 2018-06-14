package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	engine := gin.Default()

	engine.GET("ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	engine.GET("hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!!"})
	})

	engine.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
