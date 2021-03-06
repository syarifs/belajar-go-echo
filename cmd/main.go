package main

import (
	"belajar-go-echo/internal/core/service"
	"belajar-go-echo/internal/framework/database"
	"belajar-go-echo/internal/framework/repository/implementation"
	"belajar-go-echo/internal/framework/routes"
	"belajar-go-echo/internal/framework/transport/controller"
	"belajar-go-echo/internal/framework/transport/middleware"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
		db := database.InitDatabase("mysql")
		userRepo := implementation.NewUserRepository(db)
		userService := service.NewUserService(userRepo)
		userController := controller.NewUserController(userService)

		e := echo.New()
		middleware.Logging(e)
		routes.NewUserRoutes(e, userController, middleware.JWT())

		port, avail := os.LookupEnv("PORT")
		fmt.Println(port)
		if !avail {
				port = "8080"
		}

		e.Start(":" + port)
}
