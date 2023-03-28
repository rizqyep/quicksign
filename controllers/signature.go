package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/services"
)

type SignatureController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetOne(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type signatureController struct {
	service services.SignatureService
}

func NewSignatureController(service services.SignatureService) SignatureController {
	return &signatureController{
		service,
	}
}

func (controller *signatureController) Create(c *gin.Context) {
	var request domain.Signature
	c.ShouldBindJSON(&request)

	userID, _ := c.Get("x-user-id")

	convertedUserID, _ := strconv.Atoi(fmt.Sprintf("%v", userID))

	request.UserID = convertedUserID

	isValid, errors := request.ValidateRequest()

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
			"message": "Something went wrong",
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully created a new signature",
	})
}

func (controller *signatureController) GetAll(c *gin.Context) {

	userID, _ := c.Get("x-user-id")

	convertedUserID, _ := strconv.Atoi(fmt.Sprintf("%v", userID))

	result := controller.service.GetAll(convertedUserID)

	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something went wrong",
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully fetch signatures data",
	})
}

func (controller *signatureController) GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "Something went wrong",
		})
		return
	}

	var request domain.Signature
	request.ID = id
	userID, _ := c.Get("x-user-id")

	convertedUserID, _ := strconv.Atoi(fmt.Sprintf("%v", userID))
	result := controller.service.GetOne(convertedUserID, request)

	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something went wrong",
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully fetch signature data",
	})
}

func (controller *signatureController) Update(c *gin.Context) {
	var request domain.Signature
	c.ShouldBindJSON(&request)
	userID, _ := c.Get("x-user-id")

	convertedUserID, _ := strconv.Atoi(fmt.Sprintf("%v", userID))
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "Something went wrong",
		})
		return
	}
	request.ID = id

	isValid, errors := request.ValidateRequest()

	if !isValid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validation_errors": errors,
		})
		return
	}

	result := controller.service.Update(convertedUserID, request)

	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something went wrong",
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully updated signature",
	})
}

func (controller *signatureController) Delete(c *gin.Context) {
	var request domain.Signature
	c.ShouldBindJSON(&request)
	userID, _ := c.Get("x-user-id")

	convertedUserID, _ := strconv.Atoi(fmt.Sprintf("%v", userID))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error":   err,
			"message": "Something went wrong",
		})
		return
	}
	request.ID = id

	isValid, errors := request.ValidateRequest()
	if !isValid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validation_errors": errors,
		})
		return
	}

	result := controller.service.Delete(convertedUserID, request)

	if result.Error != "" {
		c.AbortWithStatusJSON(result.StatusCode, gin.H{
			"error":   result.Error,
			"message": "Something went wrong",
		})
		return
	}

	c.IndentedJSON(result.StatusCode, gin.H{
		"data":    result.Data,
		"message": "Successfully deleted signature",
	})
}
