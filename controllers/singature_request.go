package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/services"
)

type SignatureRequestController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	ApproveOrReject(c *gin.Context)
}

type signatureRequestController struct {
	service services.SignatureRequestService
}

func NewSignatureRequestController(service services.SignatureRequestService) SignatureRequestController {
	return &signatureRequestController{
		service,
	}
}

func (controller *signatureRequestController) Create(c *gin.Context) {
	var request domain.RequestSignatureRequest
	c.ShouldBindJSON(&request)

	fmt.Println(request)
	isValid, errors := request.ValidateRequest()

	username := c.Param("username")
	request.Username = username

	if !isValid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validation_errors": errors,
		})
		return
	}

	result := controller.service.Create(request)

	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something Went wrong",
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully send a signature request!",
	})
}

func (controller *signatureRequestController) GetAll(c *gin.Context) {
	userID, _ := c.Get("x-user-id")

	convertedUserID, _ := strconv.Atoi(fmt.Sprintf("%v", userID))
	result := controller.service.GetAll(convertedUserID)

	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something Went wrong",
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully fetch signature requests!",
	})
}

func (controller *signatureRequestController) GetOne(c *gin.Context) {
	userID, _ := c.Get("x-user-id")
	id, _ := strconv.Atoi(c.Param("id"))
	convertedUserID, _ := strconv.Atoi(fmt.Sprintf("%v", userID))

	var request domain.SignatureRequest
	request.ID = id
	result := controller.service.GetOne(convertedUserID, request)
	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something Went wrong",
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully fetch signature requests!",
	})
}

func (controller *signatureRequestController) ApproveOrReject(c *gin.Context) {
	userID, _ := c.Get("x-user-id")
	id, _ := strconv.Atoi(c.Param("id"))
	convertedUserID, _ := strconv.Atoi(fmt.Sprintf("%v", userID))

	var request domain.SignatureRequestApprovalRequest
	c.ShouldBindJSON(&request)
	request.ID = id

	result := controller.service.ApproveOrReject(convertedUserID, request)
	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something Went wrong",
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": result.CustomResponseMessage,
	})
}
