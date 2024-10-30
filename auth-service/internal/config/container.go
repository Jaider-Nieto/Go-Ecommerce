package config

import (
	"github.com/jaider-nieto/ecommerce-go/auth-service/internal/controller"
	"github.com/jaider-nieto/ecommerce-go/auth-service/internal/service"
)

func NewContainer () *controller.AuthController{
	authService := service.NewAuthService()
	authController := controller.NewAuthController(authService)

	return authController
}