package dto

import (
	"github.com/HarvinRaj/goldshop/internal/models"
)

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func ToUserModel(dto UserRequest) *models.Users {

	return &models.Users{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
