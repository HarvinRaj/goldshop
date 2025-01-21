package services

import (
	"github.com/HarvinRaj/goldshop/errors"
	"github.com/HarvinRaj/goldshop/internal/auth"
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
	repo       repositories.Repository
	jwtManager *auth.JWTManager
}

func NewUserService(repo repositories.Repository, jwtManager *auth.JWTManager) *UserService {
	return &UserService{
		repo:       repo,
		jwtManager: jwtManager,
	}
}

func (s *UserService) CreateUser(req *models.Users) error {

	emailExist, err := s.repo.IsEmailExist(req.Email)
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

	return s.repo.Save(req)
}

func (s *UserService) GetAllUsers() ([]*models.Users, error) {
	return s.repo.GetAllUsersList()
}

func (s *UserService) LoginAuthUser(req *models.Users) (string, error) {

	user, err := s.repo.GetUserByEmail(req)
	if err != nil {
		logger.ErrorLog.Query.Printf("GetUserByEmail Error: %v", err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		logger.ErrorLog.Error.Printf("Password does not match, %v", err)
		return "", err
	}

	token, err := s.jwtManager.GenerateToken(user)
	if err != nil {
		logger.ErrorLog.Auth.Printf("Failed to generate token, %v", err)
		return "", err
	}

	return token, nil
}
