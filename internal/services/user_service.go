package services

import (
	"github.com/A7med-Mido/educationalPlatformApi/internal/models"
	"github.com/A7med-Mido/educationalPlatformApi/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterStudent(user *models.User) error
	GetByEmail(email string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepo
}

func NewUserService() UserService {
	return &userService{repo: repositories.NewUserRepo()}
}

func (s *userService) RegisterStudent(user *models.User) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)
	user.Role = "student"
	return s.repo.Create(user)
}

func (s *userService) GetByEmail(email string) (*models.User, error) {
	return s.repo.FindByEmail(email)
}