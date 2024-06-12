package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lucas_cda/go-todo-microservices/auth"
)

func main() {
	r := gin.Default()
	api := r.Group("/api/v1/auth")
	api.GET("/health", auth.Health)

	r.Run(":8080")
}