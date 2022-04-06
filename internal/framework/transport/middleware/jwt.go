package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var SECRET_KEY = []byte("ninja-hattori")

func JWT() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: SECRET_KEY,
	})
}

func CreateToken(username, password string) (string, error) {
	expTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["password"] = password
	claims["exp"] = expTime
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSTR, err := token.SignedString(SECRET_KEY)


	return tokenSTR, err
}
