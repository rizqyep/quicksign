package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/services"
)

type ResetPasswordController interface {
	AcquireResetPasswordToken(c *gin.Context)
	ResetPassword(c *gin.Context)
}

type resetPasswordController struct {
	service services.ResetPasswordService
}

func NewResetPasswordController(service services.ResetPasswordService) ResetPasswordController {
	return &resetPasswordController{
		service,
	}
}

func (controller *resetPasswordController) AcquireResetPasswordToken(c *gin.Context) {
	var request domain.ResetPasswordToken

	c.ShouldBindJSON(&request)

	result := controller.service.AcquireResetPasswordToken(request)

	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something went wrong",
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully created reset password token",
	})

}

func (controller *resetPasswordController) ResetPassword(c *gin.Context) {
	var request domain.UpdatePasswordRequest
	request.ValidatePasswordConfirmed()
	c.ShouldBindJSON(&request)
	result := controller.service.ResetPassword(request)

	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something went wrong",
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully reset  password !",
	})
}
