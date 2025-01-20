package handlers

import (
	"github.com/HarvinRaj/goldshop/internal/dto"
	"github.com/HarvinRaj/goldshop/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate = validator.New()

type UserHandler struct {
	service services.Service
}

func NewUserHandler(service services.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) RegisterUser(c *gin.Context) {

	var req dto.UserRequest

	//Bind JSON to user struct
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot bind JSON into struct",
		})
		return
	}

	//Validate the struct
	if err := validate.Struct(&req); err != nil {
		errors := make(map[string]string)
		for _, err2 := range err.(validator.ValidationErrors) {
			errors[err2.Field()] = err2.Tag() // Field name and validation tag
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"validate error": errors,
		})
		return
	}

	user := dto.ToUserModel(req)

	if err := u.service.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}
