package main

import (
	"com.rmolcr/go-playground/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", handlers.GetUsers)
	router.Run("localhost:8080")
}
