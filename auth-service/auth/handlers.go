package auth

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthLogin(c *gin.Context) {
	var creds Creds
	if err := c.Bind(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}
	// Validar credenciales con el Servicio de Usuarios
	if !validCreds(creds.Email, creds.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := CreateJWT(creds.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func validCreds(email, password string) bool {
	res, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer([]byte(`{"email":"`+email+`", "password":"`+password+`"}`)))

	if err != nil || res.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func ValidateToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		return
	}

	// Extraer el token del encabezado
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
		return
	}

	tokenString := parts[1]
	claims := &jwt.RegisteredClaims{}

	//Verificar el token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Debes devolver tu clave secreta
		return []byte("SECRET_JWT"), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	
	// Token válido
	c.JSON(http.StatusOK, gin.H{"message": "Token is valid", "claims": claims})
}
