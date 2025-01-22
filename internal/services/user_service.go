package services

import (
	"time"

	"github.com/HarvinRaj/goldshop/errors"
	"github.com/HarvinRaj/goldshop/internal/auth"
	"github.com/HarvinRaj/goldshop/internal/models"
	"github.com/HarvinRaj/goldshop/internal/repositories"
	"github.com/HarvinRaj/goldshop/logger"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(*models.Users) (*models.Users, error)
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

func (s *UserService) CreateUser(user *models.Users) (*models.Users, error) {

	emailExist, err := s.repo.IsEmailExist(user.Email)
	if err != nil {
		logger.ErrorLog.Query.Printf("IsEmailExist error, email found: %v", err)
		return nil, err
	}

	if emailExist {
		logger.ErrorLog.Query.Printf("Email exists: %v", emailExist)
		return nil, errors.New(3000)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		logger.ErrorLog.Error.Printf("Failed to hash password: %v", err)
		return nil, err
	}

	user.PasswordHash = string(hashPassword)

	user.CreatedAt = time.Now()

	return s.repo.SaveAll(user)
}

func (s *UserService) GetAllUsers() ([]*models.Users, error) {
	return s.repo.GetAllUsersList()
}

func (s *UserService) LoginAuthUser(user *models.Users) (string, error) {

	userDB, err := s.repo.GetUserByUsername(user)
	if err != nil {
		logger.ErrorLog.Query.Printf("GetUserByUsername Error: %v", err)
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.PasswordHash), []byte(user.PasswordHash))
	if err != nil {
		logger.ErrorLog.Error.Printf("Password does not match, %v", err)
		return "", err
	}

	token, err := s.jwtManager.GenerateToken(userDB)
	if err != nil {
		logger.ErrorLog.Auth.Printf("Failed to generate token, %v", err)
		return "", err
	}

	return token, nil
}
