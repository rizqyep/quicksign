package repository

import "sync"

type repositoryPool struct {
	UserRepository
	ResetPasswordTokenRepository
}

var repositoryInstance *repositoryPool
var once sync.Once

func NewRepository() *repositoryPool {
	return &repositoryPool{
		UserRepository:               NewUserRepository(),
		ResetPasswordTokenRepository: NewResetPasswordTokenRepository(),
	}
}

func InitRepository() *repositoryPool {
	once.Do(func() {
		repositoryInstance = NewRepository()
	})
	return repositoryInstance
}
