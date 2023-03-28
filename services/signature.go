package services

import (
	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/repository"
)

type SignatureService interface {
	Create(request domain.Signature) domain.ServiceResponse
	GetAll(user_id int) domain.ServiceResponse
	GetOne(user_id int, request domain.Signature) domain.ServiceResponse
	Update(user_id int, request domain.Signature) domain.ServiceResponse
	Delete(user_id int, request domain.Signature) domain.ServiceResponse
}

type signatureService struct {
	repository repository.SignatureRepository
}

func NewSignatureService(repository repository.SignatureRepository) SignatureService {
	return &signatureService{
		repository,
	}
}

func (service *signatureService) Create(request domain.Signature) domain.ServiceResponse {
	result, err := service.repository.Create(request)

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
			"signature": result,
		},
	}
}

func (service *signatureService) GetAll(user_id int) domain.ServiceResponse {
	result, err := service.repository.GetAll(user_id)
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
			"signatures": result,
		},
	}
}

func (service *signatureService) GetOne(user_id int, request domain.Signature) domain.ServiceResponse {
	result, err := service.repository.GetOne(request)
	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	if result.UserID != user_id {
		return domain.ServiceResponse{
			Error:      "You are not authorized to access this data",
			StatusCode: 401,
			Data:       map[string]interface{}{},
		}
	}

	return domain.ServiceResponse{
		Error:      "",
		StatusCode: 201,
		Data: map[string]interface{}{
			"signature": result,
		},
	}
}

func (service *signatureService) Update(user_id int, request domain.Signature) domain.ServiceResponse {
	result, err := service.repository.GetOne(request)
	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	if result.UserID != user_id {
		return domain.ServiceResponse{
			Error:      "You are not authorized to access this data",
			StatusCode: 401,
			Data:       map[string]interface{}{},
		}
	}

	err = service.repository.Update(request)
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
	}
}

func (service *signatureService) Delete(user_id int, request domain.Signature) domain.ServiceResponse {
	result, err := service.repository.GetOne(request)
	if err != nil {
		return domain.ServiceResponse{
			Error:      err.Error(),
			StatusCode: 500,
			Data:       map[string]interface{}{},
		}
	}

	if result.UserID != user_id {
		return domain.ServiceResponse{
			Error:      "You are not authorized to access this data",
			StatusCode: 401,
			Data:       map[string]interface{}{},
		}
	}

	err = service.repository.Delete(request)
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
	}
}
