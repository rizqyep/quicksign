package controllers

import (
	"sync"
)

type controllersPool struct {
}

// var serviceInstance = services.InitServiceInstance()
var controllerInstance *controllersPool
var once sync.Once

func InitControllerInstance() *controllersPool {
	once.Do(func() {
		controllerInstance = NewControllerInstance()
	})
	return controllerInstance
}

func NewControllerInstance() *controllersPool {
	return &controllersPool{}
}
