package routers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rizqyep/quicksign/controllers"
	"github.com/rizqyep/quicksign/middleware"
)

func RouteHandlers(r *gin.Engine) {
	controllerInstance := controllers.InitControllerInstance()
	auth := r.Group("/auth")

	auth.POST("/register", controllerInstance.UserController.Register)
	auth.POST("/login", controllerInstance.UserController.LogIn)
	auth.POST("/reset-password/retrieve-token", controllerInstance.ResetPasswordController.AcquireResetPasswordToken)
	auth.POST("/reset-password", controllerInstance.ResetPasswordController.ResetPassword)

	signatures := r.Group("/signatures")

	signatures.Use(middleware.JwtAuthMiddleware())

	signatures.GET("/", controllerInstance.SignatureController.GetAll)
	signatures.POST("/", controllerInstance.SignatureController.Create)
	signatures.GET("/:id", controllerInstance.SignatureController.GetOne)
	signatures.PUT("/:id", controllerInstance.SignatureController.Update)
	signatures.DELETE("/:id", controllerInstance.SignatureController.Delete)

	r.GET("/test-token", middleware.JwtAuthMiddleware(), func(ctx *gin.Context) {
		userId, _ := ctx.Get("x-user-id")
		userName, _ := ctx.Get("x-user-username")
		userId, _ = strconv.Atoi(userId.(string))
		ctx.IndentedJSON(200, gin.H{
			"userId":   userId,
			"username": userName,
		})
	})
}
