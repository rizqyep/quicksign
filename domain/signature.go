package domain

import "time"

type Signature struct {
	ID             int       `json:"id"`
	Description    string    `json:"description"`
	SignatureToken string    `json:"signature_token"`
	QrCodeUrl      string    `json:"qr_code_url"`
	RequestID      int       `json:"request_id"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (s *Signature) ValidateRequest() (bool, map[string]string) {
	errors := make(map[string]string)
	if s.Description == "" {
		errors["description"] = "Description should not be empty"
	}

	return len(errors) == 0, errors
}
