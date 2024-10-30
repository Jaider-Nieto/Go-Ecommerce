package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jaider-nieto/ecommerce-go/auth-service/internal/config"
	"github.com/jaider-nieto/ecommerce-go/auth-service/internal/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := gin.Default()

	c := config.NewContainer()

	routes.AuthRoutes(router, c)

	router.Run(os.Getenv("PORT"))
}
