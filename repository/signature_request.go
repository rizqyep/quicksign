package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/rizqyep/quicksign/database"
	"github.com/rizqyep/quicksign/domain"
)

type SignatureRequestRepository interface {
	Create(request domain.SignatureRequest) (domain.SignatureRequest, error)
	GetAll(user_id int) ([]domain.SignatureRequest, error)
	GetOne(request domain.SignatureRequest) (domain.SignatureRequest, error)
	UpdateStatus(request domain.SignatureRequest) error
}

type signatureRequestRepository struct {
	db *sql.DB
}

func NewSignatureRequestRepository() SignatureRequestRepository {
	return &signatureRequestRepository{
		db: database.GetDBConnection(),
	}
}

func (r *signatureRequestRepository) Create(request domain.SignatureRequest) (domain.SignatureRequest, error) {
	var result domain.SignatureRequest
	sqlStatement := "INSERT INTO signature_requests (description, requester_email, requester_name, approver_id, status) VALUES ($1, $2, $3, $4, $5) RETURNING id, description, requester_email, requester_name, approver_id, status, created_at, updated_at"

	err := r.db.QueryRow(sqlStatement, request.Description, request.RequesterEmail, request.RequesterName, request.ApproverID, "PENDING").
		Scan(&result.ID, &result.Description, &result.RequesterEmail, &result.RequesterName, &result.ApproverID, &result.Status, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		return domain.SignatureRequest{}, err
	}

	return result, nil
}

func (r *signatureRequestRepository) GetAll(user_id int) ([]domain.SignatureRequest, error) {
	var results []domain.SignatureRequest

	sqlStatement := "SELECT  id, description, requester_email, requester_name, approver_id, status, created_at, updated_at FROM signature_requests WHERE approver_id =$1"
	rows, err := r.db.Query(sqlStatement, user_id)

	if err != nil {
		return results, err
	}

	defer rows.Close()

	for rows.Next() {
		var result domain.SignatureRequest

		err = rows.Scan(&result.ID, &result.Description, &result.RequesterEmail, &result.RequesterName, &result.ApproverID, &result.Status, &result.CreatedAt, &result.UpdatedAt)

		if err != nil {
			return []domain.SignatureRequest{}, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (r *signatureRequestRepository) GetOne(request domain.SignatureRequest) (domain.SignatureRequest, error) {
	var result domain.SignatureRequest
	sqlStatement := "SELECT  id, description, requester_email, requester_name, approver_id, status, created_at, updated_at FROM signature_requests WHERE id =$1"
	row := r.db.QueryRow(sqlStatement, request.ID)

	switch err := row.Scan(&result.ID, &result.Description, &result.RequesterEmail, &result.RequesterName, &result.ApproverID, &result.Status, &result.CreatedAt, &result.UpdatedAt); err {
	case sql.ErrNoRows:
		return domain.SignatureRequest{}, errors.New("Record not found")
	case nil:
		return result, nil
	default:
		return domain.SignatureRequest{}, err
	}
}

func (r *signatureRequestRepository) UpdateStatus(request domain.SignatureRequest) error {
	sqlStatement := "UPDATE signature_requests SET status=$1, updated_at=$2"
	err := r.db.QueryRow(sqlStatement, request.Status, time.Now())
	return err.Err()
}
