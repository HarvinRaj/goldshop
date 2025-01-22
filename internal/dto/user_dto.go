package dto

import (
	"github.com/HarvinRaj/goldshop/internal/models"
)

type UserRegisterRequest struct {
	UserName  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func UserRegisterToUserModel(dto *UserRegisterRequest) *models.Users {

	return &models.Users{
		UserName:  dto.UserName,
		Email:     dto.Email,
		PasswordHash:  dto.Password,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}
}

type UserLoginRequest struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func UserLoginToUserModel(dto *UserLoginRequest) *models.Users {

	return &models.Users{
		UserName: dto.UserName,
		PasswordHash: dto.Password,
	}
}
