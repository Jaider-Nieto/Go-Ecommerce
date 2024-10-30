package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaider-nieto/ecommerce-go/auth-service/internal/models"
	"github.com/jaider-nieto/ecommerce-go/auth-service/internal/service"
)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (ctr *AuthController) AuthLogin(c *gin.Context) {
	var creds models.Creds
	if err := c.Bind(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}
	// Validar credenciales con el Servicio de Usuarios
	if !ctr.service.ValidCreds(creds.Email, creds.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := ctr.service.CreateJWT(creds.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ctr *AuthController) ValidateToken(c *gin.Context) {
	if err := ctr.service.ValidateToken(c.GetHeader("Authorization")); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
	}

	// Token válido
	c.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
}
