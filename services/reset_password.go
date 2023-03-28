package services

import (
	"errors"

	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/repository"
)

type ResetPasswordService interface {
	AcquireResetPasswordToken(email string) domain.ServiceResponse
	// ResetPassword(request domain.ResetPasswordToken) domain.ServiceResponse
}

type resetPasswordService struct {
	userRepository          repository.UserRepository
	resetPasswordRepository repository.ResetPasswordTokenRepository
}

func NewResetPasswordService(userRepository repository.UserRepository, resetPasswordRepository repository.ResetPasswordTokenRepository) ResetPasswordService {
	return &resetPasswordService{
		userRepository,
		resetPasswordRepository,
	}
}

func (service *resetPasswordService) AcquireResetPasswordToken(email string) domain.ServiceResponse {
	var user domain.User
	user.Email = email

	result, err := service.userRepository.GetOne(user, "email")
	if err != nil {
		return domain.ServiceResponse{
			Error:      err,
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}
	if result == (domain.User{}) {
		return domain.ServiceResponse{
			Error:      errors.New("User Not Found"),
			StatusCode: 404,
			Data:       map[string]interface{}{},
		}
	}

	request := domain.ResetPasswordToken{
		Email: email,
	}
	tokenResult, err := service.resetPasswordRepository.Create(request)

	if err != nil {
		return domain.ServiceResponse{
			Error:      err,
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	return domain.ServiceResponse{
		Error:      nil,
		StatusCode: 201,
		Data: map[string]interface{}{
			"reset_password_token": tokenResult,
		},
	}
}

func (service *resetPasswordService) ResetPassword(request domain.UpdatePasswordRequest) domain.ServiceResponse {
	resetPasswordToken := request.ResetPasswordToken
	result, err := service.resetPasswordRepository.GetOne(resetPasswordToken)
	if err != nil {
		return domain.ServiceResponse{
			Error:      err,
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}
	if result == (domain.ResetPasswordToken{}) {
		return domain.ServiceResponse{
			Error:      errors.New("Invalid token"),
			StatusCode: 400,
			Data:       map[string]interface{}{},
		}
	}
	user := domain.User{Email: request.Email, Password: request.NewPassword}
	err = service.userRepository.UpdatePassword(user)

	if err != nil {
		return domain.ServiceResponse{
			Error:      err,
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	return domain.ServiceResponse{
		Error:      nil,
		StatusCode: 200,
		Data:       map[string]interface{}{},
	}
}
