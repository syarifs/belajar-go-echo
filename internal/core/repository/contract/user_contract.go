package contract

import "belajar-go-echo/internal/core/models"

type UserRepository interface {
	GetAll() (*[]models.User, error)
	Create(models.User) error
	Login(models.User) error
}
