package handlers

import (
	"github.com/HarvinRaj/goldshop/internal/dto"
	"github.com/HarvinRaj/goldshop/internal/services"
	"github.com/HarvinRaj/goldshop/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service services.Service
}

func NewUserHandler(service services.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {

	var req *dto.UserRegisterRequest

	//Bind JSON to user struct
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot bind JSON into struct",
			"err":   err,
		})
		return
	}

	//Validate the struct
	errors, err := util.ValidateDTO(req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"errors": "validate error",
			"err":    errors,
		})
		return
	}

	user := dto.UserRegisterToUserModel(req)

	if err = h.service.CreateUser(user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "create user error",
			"err":   err,
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func (h *UserHandler) GetUsersList(c *gin.Context) {

	users, err := h.service.GetAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "get all user list error",
			"err":   err,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *UserHandler) Login(c *gin.Context) {

	var req *dto.UserLoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot bind JSON into struct",
			"err":   err,
		})
		return
	}

	errors, err := util.ValidateDTO(req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "validate dto error",
			"err":   errors,
		})
		return
	}

	user := dto.UserLoginToUserModel(req)

	token, err := h.service.LoginAuthUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error": "login authenthication error",
			"err":   err,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"token": token,
	})

}
