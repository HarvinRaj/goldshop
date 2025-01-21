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

func (u *UserHandler) RegisterUser(c *gin.Context) {

	var req *dto.UserRegisterRequest

	//Bind JSON to user struct
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot bind JSON into struct",
		})
		return
	}

	//Validate the struct
	errors, err := util.ValidateDTO(&req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"validate errors": errors,
		})
	}

	user := dto.UserRegisterToUserModel(req)

	if err = u.service.CreateUser(user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func (u *UserHandler) GetUsersList(c *gin.Context) {

	users, err := u.service.GetAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (u *UserHandler) Login(c *gin.Context) {

	var req *dto.UserLoginRequest

	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": "Cannot bind JSON into struct",
		})
		return
	}

	errors, err := util.ValidateDTO(&req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"validate error": errors,
		})
	}

	user := dto.UserLoginToUserModel(req)

	token, err := u.service.LoginAuthUser(user)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"Authorization": err,
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"token": token,
	})

}
