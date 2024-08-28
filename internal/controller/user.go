package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vegetarian23/go-initializr/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": uc.userService.GetUserById(),
		"users": []string{"vegetarian23", "vegetarian24"},
	})
}
