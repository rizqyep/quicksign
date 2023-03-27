package repository

import "sync"

type repositoryPool struct {
}

var repositoryInstance *repositoryPool
var once *sync.Once

func NewRepository() *repositoryPool {
	return &repositoryPool{}
}

func InitRepository() *repositoryPool {
	once.Do(func() {
		repositoryInstance = NewRepository()
	})
	return repositoryInstance
}
