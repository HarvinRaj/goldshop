package services

import (
	"github.com/HarvinRaj/goldshop/errors"
	"github.com/HarvinRaj/goldshop/internal/models"
	"github.com/HarvinRaj/goldshop/internal/repositories"
	"github.com/HarvinRaj/goldshop/logger"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(user *models.Users) error
}

type UserService struct {
	repo repositories.Repository
}

func NewUserService(repo repositories.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(user *models.Users) error {

	emailExist, err := u.repo.IsEmailExist(user.Email)
	if err != nil {
		logger.ErrorLog.Query.Printf("IsEmailExist error, %v", err)
		return err
	}

	if emailExist {
		logger.ErrorLog.Query.Printf("Email exists: %v", emailExist)
		return errors.New(3000)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLog.Error.Printf("Failed to hash password: %v", err)
		return err
	}

	user.Password = string(hashPassword)

	return u.repo.Save(user)
}
