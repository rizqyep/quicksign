package domain

import (
	"net/mail"
	"time"
)

type SignatureRequest struct {
	ID             int       `json:"id"`
	RequesterEmail string    `json:"requester_email"`
	RequesterName  string    `json:"requester_name"`
	Description    string    `json:"requester_description"`
	ApproverID     int       `json:"approver_id"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type RequestSignatureRequest struct {
	SignatureRequest
	Username string `json:"username"`
}

type SignatureRequestApprovalRequest struct {
	SignatureRequest
	OverrideDescription bool   `json:"override_description"`
	NewDescription      string `json:"new_description"`
}

func (r *RequestSignatureRequest) ValidateRequest() (bool, map[string]string) {
	errors := make(map[string]string)

	if r.Description == "" {
		errors["description"] = "Description should not be empty"
	}

	_, err := mail.ParseAddress(r.RequesterEmail)

	if err != nil {
		errors["email"] = "Email should be a valid email format!"
	}

	if r.RequesterName == "" {
		errors["requester_name"] = "Requester name should not be empty!"
	}

	return len(errors) == 0, errors
}

func (r *SignatureRequestApprovalRequest) ValidateRequest() (bool, map[string]string) {
	errors := make(map[string]string)

	if r.OverrideDescription == true {
		if r.NewDescription == "" {
			errors["new_description"] = "New description should be filled!"
		}
	}
	return len(errors) == 0, errors
}
