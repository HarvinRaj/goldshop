package users

import "github.com/gin-gonic/gin"

type UserHandler struct {
	service Service
}

func NewUserHandler(service Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) CreateUser(c *gin.Context) {

}
