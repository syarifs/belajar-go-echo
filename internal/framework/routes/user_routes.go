package routes

import (
	"belajar-go-echo/internal/framework/transport/controller"

	"github.com/labstack/echo/v4"
)

func NewUserRoutes(e *echo.Echo, ucon *controller.UserController, middleware ...echo.MiddlewareFunc) {
	e.GET("/user", ucon.GetAll, middleware...)
	e.POST("/user", ucon.Create)
	e.POST("/login", ucon.Login)
}
