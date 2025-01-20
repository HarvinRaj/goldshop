package dto

import (
	"github.com/HarvinRaj/goldshop/internal/models"
)

type UserRegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func UserRegisterToUserModel(dto *UserRegisterRequest) *models.Users {

	return &models.Users{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func UserLoginToUserModel(dto *UserLoginRequest) *models.Users {

	return &models.Users{
		Email:    dto.Email,
		Password: dto.Password,
	}
}
