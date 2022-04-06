package service

import (
	"belajar-go-echo/internal/core/models"
	"belajar-go-echo/internal/core/repository/contract"
)

type UserService struct {
	Repo contract.UserRepository
}

func NewUserService(repo contract.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (serv UserService) GetAllUser() (user *[]models.User, err error) {
	user, err = serv.Repo.GetAll()
	return
}

func (serv UserService) CreateUser(user models.User) (err error) {
	err = serv.Repo.Create(user)
	return
}

func (serv UserService) Login(user models.User) (err error) {
	err = serv.Repo.Login(user)
	return
	
}
