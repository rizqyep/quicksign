package services

import (
	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/repository"
	"github.com/rizqyep/quicksign/utils"
)

type ResetPasswordService interface {
	AcquireResetPasswordToken(request domain.ResetPasswordToken) domain.ServiceResponse
	ResetPassword(request domain.UpdatePasswordRequest) domain.ServiceResponse
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

func (service *resetPasswordService) AcquireResetPasswordToken(request domain.ResetPasswordToken) domain.ServiceResponse {
	var user domain.User
	user.Email = request.Email

	result, err := service.userRepository.GetOne(user, "email")
	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}
	if result == (domain.User{}) {
		return domain.ServiceResponse{
			Error:      "User with given email does not exist",
			StatusCode: 400,
			Data:       map[string]interface{}{},
		}
	}

	tokenResult, err := service.resetPasswordRepository.Create(request)

	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	payload := utils.ResetPasswordMailPayload{
		Email: request.Email,
		Token: tokenResult.Token,
	}

	//Run in go routine to make email sending run in background (not blocking http process)
	go utils.SendResetPasswordLink(payload)

	return domain.ServiceResponse{
		Error:      "",
		StatusCode: 201,
		Data: map[string]interface{}{
			"reset_password_token": tokenResult,
		},
	}
}

func (service *resetPasswordService) ResetPassword(request domain.UpdatePasswordRequest) domain.ServiceResponse {
	resetPasswordToken := request.ResetPasswordToken
	result, err := service.resetPasswordRepository.GetOne(resetPasswordToken)
	if result == (domain.ResetPasswordToken{}) {
		return domain.ServiceResponse{
			Error:      "Invalid Token",
			StatusCode: 400,
			Data:       map[string]interface{}{},
		}
	}
	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	user := domain.User{Email: result.Email, Password: request.NewPassword}
	err = service.userRepository.UpdatePassword(user)

	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	err = service.resetPasswordRepository.Invalidate(result)
	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}
	return domain.ServiceResponse{
		Error:      "",
		StatusCode: 200,
		Data:       map[string]interface{}{},
	}
}
