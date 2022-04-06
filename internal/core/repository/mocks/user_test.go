package mocks

import (
	"belajar-go-echo/internal/core/models"
	"belajar-go-echo/internal/core/service"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_GetAllUser(t *testing.T) {
	mockUser := new(UserRepository)
	t.Run("success", func(t *testing.T) {
		mockUser.On("GetAll").Return(&[]models.User{}, nil).Once()
		userService := service.NewUserService(mockUser)
		user, err := userService.Repo.GetAll()
		assert.NotNil(t, user)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		mockUser.On("GetAll").Return(nil, errors.New("fail to mock GetAll")).Once()
		userService := service.NewUserService(mockUser)
		user, err := userService.Repo.GetAll()
		assert.Nil(t, user)
		assert.Error(t, err)
	})
}

func Test_CreateUser(t *testing.T) {
	mockUser := new(UserRepository)
	mockData := models.User{
		Email: "test@mock.com",
	Password: "password",
	}
	
	t.Run("success", func(t *testing.T) {
		mockUser.On("Create", mock.Anything).Return(nil).Once()
		userService := service.NewUserService(mockUser)
		err := userService.Repo.Create(mockData)
		assert.NoError(t, err)
	})
	
	t.Run("fail", func(t *testing.T) {
		mockUser.On("Create", mock.Anything).Return(errors.New("mock create fail")).Once()
		userService := service.NewUserService(mockUser)
		err := userService.Repo.Create(mockData)
		assert.Error(t, err)
	})
}

func Test_Login(t *testing.T) {
	mockUser := new(UserRepository)
	mockData := models.User{
		Email: "test2@mock.com",
		Password: "password",
	}
	
	t.Run("success", func(t *testing.T) {
		mockUser.On("Login", mock.Anything).Return(nil).Once()
		userService := service.NewUserService(mockUser)
		err := userService.Repo.Login(mockData)
		assert.NoError(t, err)
	})
	
	t.Run("fail", func(t *testing.T) {
		mockUser.On("Login", mock.Anything).Return(errors.New("mock create fail")).Once()
		userService := service.NewUserService(mockUser)
		err := userService.Repo.Login(mockData)
		assert.Error(t, err)
	})
}
