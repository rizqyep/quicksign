package repository

import (
	"database/sql"
	"errors"

	"github.com/rizqyep/quicksign/database"
	"github.com/rizqyep/quicksign/domain"
	"github.com/rizqyep/quicksign/utils"
)

type SignatureRepository interface {
	Create(request domain.Signature) (domain.Signature, error)
	GetAll(user_id int) ([]domain.Signature, error)
	GetOne(request domain.Signature) (domain.Signature, error)
	Update(request domain.Signature) error
	Delete(request domain.Signature) error
}

type signatureRepository struct {
	db *sql.DB
}

func NewSignatureRepository() SignatureRepository {
	return &signatureRepository{
		db: database.GetDBConnection(),
	}
}

func (r *signatureRepository) Create(request domain.Signature) (domain.Signature, error) {
	var result domain.Signature
	request.SignatureToken = utils.RandStringRunes(64)
	request.QrCodeUrl = utils.ConstructSignatureQRCode(request.SignatureToken)

	sqlStatement := "INSERT INTO signatures (signature_token, qr_code_url, description, request_id, user_id) VALUES ($1, $2, $3 ,$4, $5) RETURNING id, signature_token, qr_code_url, description, COALESCE(request_id, 0) request_id, user_id, created_at, updated_at"

	err := r.db.QueryRow(sqlStatement, &request.SignatureToken, &request.QrCodeUrl, &request.Description, nil, &request.UserID).Scan(&result.ID, &result.SignatureToken, &result.QrCodeUrl, &result.Description, &result.RequestID, &result.UserID, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		return domain.Signature{}, err
	}

	return result, nil
}

func (r *signatureRepository) GetAll(user_id int) ([]domain.Signature, error) {
	var results []domain.Signature
	sqlStatement := "SELECT id, signature_token, qr_code_url, description,COALESCE(request_id, 0) request_id, user_id, created_at, updated_at from signatures WHERE user_id = $1"

	rows, err := r.db.Query(sqlStatement, user_id)

	if err != nil {
		return results, err
	}

	defer rows.Close()

	for rows.Next() {
		var result domain.Signature
		err = rows.Scan(&result.ID, &result.SignatureToken, &result.QrCodeUrl, &result.Description, &result.RequestID, &result.UserID, &result.CreatedAt, &result.UpdatedAt)

		if err != nil {
			return []domain.Signature{}, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *signatureRepository) GetOne(request domain.Signature) (domain.Signature, error) {
	var result domain.Signature
	sqlStatement := "SELECT id, signature_token, qr_code_url, description, COALESCE(request_id, 0) request_id, user_id, created_at, updated_at from signatures WHERE id = $1"

	row := r.db.QueryRow(sqlStatement, request.ID)

	switch err := row.Scan(&result.ID, &result.SignatureToken, &result.QrCodeUrl, &result.Description, &result.RequestID, &result.UserID, &result.CreatedAt, &result.UpdatedAt); err {
	case sql.ErrNoRows:
		return domain.Signature{}, errors.New("Record not found")
	case nil:
		return result, nil
	default:
		return domain.Signature{}, err
	}

}

func (r *signatureRepository) Update(request domain.Signature) error {
	sqlStatement := "UPDATE signatures SET description=$1 WHERE id = $2"

	err := r.db.QueryRow(sqlStatement, &request.Description, &request.ID)

	return err.Err()
}

func (r *signatureRepository) Delete(request domain.Signature) error {
	sqlStatement := "DELETE FROM signatures WHERE id = $1"
	err := r.db.QueryRow(sqlStatement, &request.ID)

	return err.Err()
}
