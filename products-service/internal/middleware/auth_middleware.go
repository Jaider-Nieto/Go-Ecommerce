package middleware

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware es el middleware para validar el token JWT
func AuthMiddleware(c *gin.Context) {
	// Obtener el token del encabezado Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.Abort()
		return
	}

	// Extraer el token del encabezado
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
		c.Abort()
		return
	}

	tokenString := parts[1]

	// Crear solicitud para validar el token
	req, err := http.NewRequest("GET", "http://172.17.0.1:8081/validate-token", nil)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to create request"})
		c.Abort()
		return
	}
	req.Header.Set("Authorization", "Bearer "+tokenString)

	// Enviar la solicitud y obtener la respuesta
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to validate token: " + err.Error()})
		c.Abort()
		return
	}
	defer res.Body.Close()

	// Leer el cuerpo de la respuesta para depuración si es necesario
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		c.Abort()
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token", "details": string(body)})
		c.Abort()
		return
	}

	// Si el token es válido, continuar con la solicitud
	c.Next()
}
