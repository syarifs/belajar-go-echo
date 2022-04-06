package controller

import (
	"belajar-go-echo/internal/core/models"
	"belajar-go-echo/internal/core/service"
	"belajar-go-echo/internal/framework/transport/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	srv *service.UserService
}

func NewUserController(srv *service.UserService) *UserController {
	return &UserController{srv}
}

func (ucon UserController) GetAll(c echo.Context) error {
	user, err := ucon.srv.GetAllUser()
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Data Fetched",
			"data": user,
		})
	}
}

func (ucon UserController) Create(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	err := ucon.srv.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	} else {
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "Data Saved",
			"data": user,
		})
	}
}

func (ucon UserController) Login(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	err := ucon.srv.Login(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	} else {
		token, _ := middleware.CreateToken(user.Email, user.Password)
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Logged In",
			"data": user,
			"token": token,
		})
	}
}
