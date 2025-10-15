package repositories

import (
	"github.com/A7med-Mido/educationalPlatformApi/internal/database"
	"github.com/A7med-Mido/educationalPlatformApi/internal/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo() UserRepo {
	return &userRepo{db: database.DB}
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}