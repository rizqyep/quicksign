package domain

import "time"

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
	SignatureRequest SignatureRequest
	Username         string
}

type SingatureRequestApprovalRequest struct {
	Signature           Signature
	SignatureRequest    SignatureRequest
	OverrideDescription bool
	NewDescription      string
}
