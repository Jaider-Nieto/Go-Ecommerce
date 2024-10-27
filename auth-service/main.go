package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaider-nieto/ecommerce-go/auth-service/auth"
)

func main() {
	router := gin.Default()

	router.POST("/auth", auth.AuthLogin)
	router.GET("/validate-token", auth.ValidateToken)

	router.Run(":8081")
}
