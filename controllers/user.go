package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/services"
	"github.com/rizqyep/quicksign/validations"
)

type UserController interface {
	Register(c *gin.Context)
	LogIn(c *gin.Context)
}

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{
		service,
	}
}

func (controller *userController) Register(c *gin.Context) {
	var user domain.User

	c.ShouldBindJSON(&user)

	isRequestValid, errors := validations.ValidateRegistrationRequest(user)

	if !isRequestValid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":           "Some fields are incorrect",
			"validation_errors": errors,
		})
		return
	}

	result := controller.service.Register(user)

	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something went wrong",
			"data":    nil,
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully Registered!",
	})
}

func (controller *userController) LogIn(c *gin.Context) {
	var user domain.User

	c.ShouldBindJSON(&user)

	isRequestValid, errors := validations.ValidateLoginRequest(user)

	if !isRequestValid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":           "Some fields are incorrect",
			"validation_errors": errors,
		})
		return
	}

	result := controller.service.LogIn(user)

	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something went wrong",
			"data":    nil,
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully Registered!",
	})
}
