package services

import (
	"sync"

	"github.com/rizqyep/quicksign/repository"
)

type servicesPool struct {
}

var repositoryInstance = repository.NewRepository()
var serviceInstance *servicesPool
var once sync.Once

func InitServiceInstance() *servicesPool {
	once.Do(func() {
		serviceInstance = NewServiceInstance()
	})
	return serviceInstance
}

func NewServiceInstance() *servicesPool {
	return &servicesPool{}
}
