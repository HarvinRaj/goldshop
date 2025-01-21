package services

import (
	"github.com/HarvinRaj/goldshop/errors"
	"github.com/HarvinRaj/goldshop/internal/models"
	"github.com/HarvinRaj/goldshop/internal/repositories"
	"github.com/HarvinRaj/goldshop/logger"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(*models.Users) error
	GetAllUsers() ([]*models.Users, error)
	LoginAuthUser(*models.Users) (string, error)
}

type UserService struct {
	repo repositories.Repository
}

func NewUserService(repo repositories.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(req *models.Users) error {

	emailExist, err := u.repo.IsEmailExist(req.Email)
	if err != nil {
		logger.ErrorLog.Query.Printf("IsEmailExist error, email found: %v", err)
		return err
	}

	if emailExist {
		logger.ErrorLog.Query.Printf("Email exists: %v", emailExist)
		return errors.New(3000)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLog.Error.Printf("Failed to hash password: %v", err)
		return err
	}

	req.Password = string(hashPassword)

	return u.repo.Save(req)
}

func (u *UserService) GetAllUsers() ([]*models.Users, error) {
	return u.repo.GetAllUsersList()
}

func (u *UserService) LoginAuthUser(req *models.Users) (string, error) {

	loginUser, err := u.repo.GetUserByEmail(req)
	if err != nil {
		logger.ErrorLog.Query.Printf("GetUserByEmail Error: %v", err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(req.Password))
	if err != nil {
		logger.ErrorLog.Error.Printf("Password does not match, %v", err)
		return "", err
	}

	return "", nil
}
