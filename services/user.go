package services

import (
	"fmt"

	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/repository"
	"github.com/rizqyep/quicksign/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(domain.User) domain.ServiceResponse
	LogIn(domain.User) domain.ServiceResponse
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository,
	}
}

func (service *userService) Register(user domain.User) domain.ServiceResponse {
	err, result := service.repository.Create(user)

	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	token, err := utils.CreateJWTToken(&result)
	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	return domain.ServiceResponse{
		Error:      "",
		StatusCode: 201,
		Data: map[string]interface{}{
			"user":  result,
			"token": token,
		},
	}

}

func (service *userService) LogIn(user domain.User) domain.ServiceResponse {

	result, err := service.repository.GetOneForAuth(user)

	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}
	if result == (domain.User{}) {
		return domain.ServiceResponse{
			Error:      "Invalid Credentials",
			StatusCode: 400,
			Data:       map[string]interface{}{},
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		fmt.Println(err)
		return domain.ServiceResponse{
			Error:      "Invalid Credentials",
			StatusCode: 401,
			Data:       map[string]interface{}{},
		}
	}
	token, err := utils.CreateJWTToken(&result)

	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}
	result.HidePassword()
	return domain.ServiceResponse{
		Error:      "",
		StatusCode: 201,
		Data: map[string]interface{}{
			"user":  result,
			"token": token,
		},
	}

}
