package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rizqyep/quicksign/controllers"
	"github.com/rizqyep/quicksign/middleware"
)

func RouteHandlers(r *gin.Engine) {
	controllerInstance := controllers.InitControllerInstance()

	r.POST("/register", controllerInstance.UserController.Register)
	r.POST("/login", controllerInstance.UserController.LogIn)

	r.GET("/test-token", middleware.JwtAuthMiddleware(), func(ctx *gin.Context) {
		userId, _ := ctx.Get("x-user-id")
		userName, _ := ctx.Get("x-user-username")

		ctx.IndentedJSON(200, gin.H{
			"userId":   userId,
			"username": userName,
		})
	})
}
