package domain

import "time"

type ResetPasswordToken struct {
	ID        int       `json:"id"`
	Token     string    `json:"token"`
	Email     string    `json:"email"`
	Valid     bool      `json:"valid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdatePasswordRequest struct {
	ResetPasswordToken
	NewPassword        string `json:"new_password"`
	ConfirmNewPassword string `json:"confirm_new_password"`
}

func (updateRequest *UpdatePasswordRequest) ValidatePasswordConfirmed() bool {
	return updateRequest.NewPassword == updateRequest.ConfirmNewPassword
}
