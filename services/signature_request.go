package services

import (
	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/repository"
	"github.com/rizqyep/quicksign/utils"
)

type SignatureRequestService interface {
	Create(request domain.RequestSignatureRequest) domain.ServiceResponse
	GetAll(user_id int) domain.ServiceResponse
	GetOne(user_id int, request domain.SignatureRequest) domain.ServiceResponse
	ApproveOrReject(user_id int, request domain.SignatureRequestApprovalRequest) domain.ServiceResponse
}

type signatureRequestService struct {
	repository          repository.SignatureRequestRepository
	signatureRepository repository.SignatureRepository
	userRepository      repository.UserRepository
}

func NewSignatureRequestService(repository repository.SignatureRequestRepository, signatureRepository repository.SignatureRepository, userRepository repository.UserRepository) SignatureRequestService {
	return &signatureRequestService{
		repository,
		signatureRepository,
		userRepository,
	}
}

func (service *signatureRequestService) Create(request domain.RequestSignatureRequest) domain.ServiceResponse {
	var user domain.User
	user.Username = request.Username
	userResult, err := service.userRepository.GetOne(user, "username")

	if userResult == (domain.User{}) {
		return domain.ServiceResponse{
			StatusCode: 404,
			Error:      "No user found with that username",
			Data:       map[string]interface{}{},
		}
	}

	if err != nil {
		return domain.ServiceResponse{
			StatusCode: 500,
			Error:      err.Error(),
			Data:       map[string]interface{}{},
		}
	}
	request.ApproverID = userResult.ID
	result, err := service.repository.Create(request.SignatureRequest)

	if err != nil {
		return domain.ServiceResponse{
			StatusCode: 500,
			Error:      err.Error(),
			Data:       map[string]interface{}{},
		}
	}

	return domain.ServiceResponse{
		StatusCode: 201,
		Error:      "",
		Data: map[string]interface{}{
			"signature_request": result,
		},
	}
}

func (service *signatureRequestService) GetAll(user_id int) domain.ServiceResponse {
	result, err := service.repository.GetAll(user_id)
	if err != nil {
		return domain.ServiceResponse{
			StatusCode: 500,
			Error:      err.Error(),
			Data:       map[string]interface{}{},
		}
	}

	return domain.ServiceResponse{
		StatusCode: 201,
		Error:      "",
		Data: map[string]interface{}{
			"signature_requests": result,
		},
	}

}

func (service *signatureRequestService) GetOne(user_id int, request domain.SignatureRequest) domain.ServiceResponse {
	result, err := service.repository.GetOne(request)
	if err != nil {
		return domain.ServiceResponse{
			StatusCode: 500,
			Error:      err.Error(),
			Data:       map[string]interface{}{},
		}
	}
	if result.ApproverID != user_id {
		return domain.ServiceResponse{
			StatusCode: 401,
			Error:      "You are not allowed to access this data",
			Data:       map[string]interface{}{},
		}
	}
	return domain.ServiceResponse{
		StatusCode: 201,
		Error:      "",
		Data: map[string]interface{}{
			"signature_request": result,
		},
	}

}

func (service *signatureRequestService) ApproveOrReject(user_id int, request domain.SignatureRequestApprovalRequest) domain.ServiceResponse {

	result, err := service.repository.GetOne(request.SignatureRequest)
	if err != nil {
		return domain.ServiceResponse{
			StatusCode: 500,
			Error:      err.Error(),
			Data:       map[string]interface{}{},
		}
	}
	if result.ApproverID != user_id {
		return domain.ServiceResponse{
			StatusCode: 401,
			Error:      "You are not allowed to access this data",
			Data:       map[string]interface{}{},
		}
	}

	if result.Status != "PENDING" {
		return domain.ServiceResponse{
			StatusCode: 400,
			Error:      "You already approved / rejected this application previously",
			Data:       map[string]interface{}{},
		}
	}
	// Additional Action for Approved Request
	if request.Status == "APPROVED" {
		var newSignatureRequest domain.Signature
		if request.OverrideDescription {
			newSignatureRequest.Description = request.NewDescription
		} else {
			newSignatureRequest.Description = result.Description
		}
		signature, err := service.signatureRepository.Create(newSignatureRequest)

		if err != nil {
			return domain.ServiceResponse{
				StatusCode: 500,
				Error:      err.Error(),
				Data:       map[string]interface{}{},
			}
		}

		mailPayload := utils.SignatureMailPayload{
			QrCodeUrl:      signature.QrCodeUrl,
			RequesterEmail: result.RequesterEmail,
			RequesterName:  result.RequesterName,
		}

		go utils.SendSignatureMail(mailPayload)
		return domain.ServiceResponse{
			Error: "",
			Data: map[string]interface{}{
				"new_signature": signature,
			},
			CustomResponseMessage: "Successfully Approved a Signature Request and an E-mail with attached signature has been sent to signature's requester!",
		}
	}

	//Inform rejected via email to requester
	mailPayload := utils.RejectedSignatureMailPayload{
		RequesterEmail: result.RequesterEmail,
		RequesterName:  result.RequesterName,
		Description:    result.Description,
	}
	//Call with go routine to un-block request process
	go utils.SendRejectedSignatureMail(mailPayload)

	// Update Request status
	err = service.repository.UpdateStatus(request.SignatureRequest)
	if err != nil {
		return domain.ServiceResponse{
			StatusCode: 500,
			Error:      err.Error(),
			Data:       map[string]interface{}{},
		}
	}

	return domain.ServiceResponse{
		Error:                 "",
		Data:                  map[string]interface{}{},
		CustomResponseMessage: "Successfully Rejected Signature Request",
	}
}
