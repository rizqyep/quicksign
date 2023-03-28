package services

import (
	"sync"

	"github.com/rizqyep/quicksign/repository"
)

type servicesPool struct {
	UserService
	ResetPasswordService
}

var repositoryInstance = repository.InitRepository()
var serviceInstance *servicesPool
var once sync.Once

func InitServiceInstance() *servicesPool {
	once.Do(func() {
		serviceInstance = NewServiceInstance()
	})
	return serviceInstance
}

func NewServiceInstance() *servicesPool {
	return &servicesPool{
		UserService:          NewUserService(repositoryInstance.UserRepository),
		ResetPasswordService: NewResetPasswordService(repositoryInstance.UserRepository, repositoryInstance.ResetPasswordTokenRepository),
	}
}
