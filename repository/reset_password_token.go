package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rizqyep/quicksign/database"
	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/utils"
)

type ResetPasswordTokenRepository interface {
	Create(request domain.ResetPasswordToken) (domain.ResetPasswordToken, error)
	GetOne(request domain.ResetPasswordToken) (domain.ResetPasswordToken, error)
	Invalidate(request domain.ResetPasswordToken) error
}

type resetPasswordTokenRepository struct {
	db *sql.DB
}

func NewResetPasswordTokenRepository() ResetPasswordTokenRepository {
	return &resetPasswordTokenRepository{
		db: database.GetDBConnection(),
	}
}

func (r *resetPasswordTokenRepository) Create(request domain.ResetPasswordToken) (domain.ResetPasswordToken, error) {
	var result domain.ResetPasswordToken
	request.Token = utils.RandStringRunes(32)
	sqlStatement := "INSERT INTO reset_password_tokens (token, email) VALUES ($1, $2) RETURNING id, email, token, valid, created_at, updated_at"
	err := r.db.QueryRow(sqlStatement, request.Token, request.Email).Scan(&result.ID, &result.Email, &result.Token, &result.Valid, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return domain.ResetPasswordToken{}, err
	}
	return result, nil
}

func (r *resetPasswordTokenRepository) GetOne(request domain.ResetPasswordToken) (domain.ResetPasswordToken, error) {
	var result domain.ResetPasswordToken
	sqlStatement := "SELECT id, email, token, valid, created_at, updated_at FROM reset_password_tokens WHERE token = $1 AND valid = $2"

	row := r.db.QueryRow(sqlStatement, request.Token, true)
	switch err := row.Scan(&result.ID, &result.Email, &result.Token, &result.Valid, &result.CreatedAt, &result.UpdatedAt); err {
	case sql.ErrNoRows:
		return domain.ResetPasswordToken{}, errors.New(fmt.Sprintf("record not found"))
	case nil:
		return result, nil
	default:
		return domain.ResetPasswordToken{}, err
	}
}

func (r *resetPasswordTokenRepository) Invalidate(request domain.ResetPasswordToken) error {
	sqlStatement := "UPDATE reset_password_tokens SET valid=$1 WHERE id=$2"
	err := r.db.QueryRow(sqlStatement, false, request.ID)

	return err.Err()
}
