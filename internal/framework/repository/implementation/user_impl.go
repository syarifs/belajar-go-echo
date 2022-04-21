package implementation

import (
	"belajar-go-echo/internal/core/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (repo userRepository) GetAll() (user *[]models.User, err error) {
	err = repo.db.Find(&user).Error
	return
}

func (repo userRepository) Create(user models.User) (err error) {
	err = repo.db.Save(&user).Error
	return
}

func (repo userRepository) Login(user models.User) (err error) {
	err = repo.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&models.User{}).Error
	return
}

