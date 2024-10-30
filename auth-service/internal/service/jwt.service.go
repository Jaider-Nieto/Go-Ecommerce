package service

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"sub": email,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("SECRET_JWT"))
}

func (s *AuthService) ValidCreds(email, password string) bool {
	res, err := http.Post("http://localhost:8080/login", "application/json", bytes.NewBuffer([]byte(`{"email":"`+email+`", "password":"`+password+`"}`)))

	if err != nil || res.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func (s *AuthService) ValidateToken(authHeader string) error {
	if authHeader == "" {
		return errors.New("authorization header is missing")
	}

	// Extraer el token del encabezado
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return errors.New("invalid authorization format")
	}

	tokenString := parts[1]
	claims := &jwt.RegisteredClaims{}

	//Verificar el token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Debes devolver tu clave secreta
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil || !token.Valid {
		return errors.New("invalid token")
	}
	return nil
}
