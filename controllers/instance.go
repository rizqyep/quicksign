package controllers

import (
	"sync"

	"github.com/rizqyep/quicksign/services"
)

type controllersPool struct {
	UserController
}

var serviceInstance = services.InitServiceInstance()
var controllerInstance *controllersPool
var once sync.Once

func InitControllerInstance() *controllersPool {
	once.Do(func() {
		controllerInstance = NewControllerInstance()
	})
	return controllerInstance
}

func NewControllerInstance() *controllersPool {
	return &controllersPool{
		UserController: NewUserController(serviceInstance.UserService),
	}
}
