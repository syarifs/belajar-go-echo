package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var SECRET_KEY = []byte("864b1bd71cfb9d62e2645bfda9a3bb98643bb9b7d1733fad40abed2e59940083")

func JWT() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: SECRET_KEY,
	})
}

func CreateToken(username string) (tokenSTR string, err error) {
	expTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["username"] = username
	claims["exp"] = expTime
	claims["iat"] = time.Now().Unix()
	claims["nbf"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSTR, err = token.SignedString(SECRET_KEY)

	return
}
