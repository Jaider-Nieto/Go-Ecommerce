package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jaider-nieto/ecommerce-go/auth-service/internal/controller"
)

func AuthRoutes(router *gin.Engine, ctr *controller.AuthController) {
	authGroup := router.Group("auth")
	{

		authGroup.POST("/auth", ctr.AuthLogin)
		authGroup.GET("/validate-token", ctr.ValidateToken)
	}
}
