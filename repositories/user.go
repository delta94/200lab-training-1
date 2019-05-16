package repositories

import (
	"github.com/200lab-training-1/models"
	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	Create(models.User) (*models.User, error)
	Find(string) (*models.User, error)
}

type UserRepoImpl struct {
	DB *gorm.DB
}

func (userRepo *UserRepoImpl) Create(user models.User) (*models.User, error) {
	err := userRepo.DB.Create(&user).Error
	return &user, err
}

func (userRepo *UserRepoImpl) Find(email string) (*models.User, error) {
	user := &models.User{}
	err := userRepo.DB.Where("email = ?", email).First(user).Error
	return user, err
}
