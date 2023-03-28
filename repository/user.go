package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/rizqyep/quicksign/database"
	"github.com/rizqyep/quicksign/domain"
)

type UserRepository interface {
	Create(user domain.User) (error, domain.User)
	GetOne(user domain.User, option string) (domain.User, error)
	GetOneForAuth(user domain.User) (domain.User, error)
	UpdatePassword(user domain.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: database.GetDBConnection(),
	}
}

func (r *userRepository) Create(user domain.User) (error, domain.User) {
	var errs error
	var result = domain.User{}
	sqlStatement := `INSERT INTO users (first_name, last_name, username, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id, first_name, last_name, username, email, created_at, updated_at`
	user.HashPassword()
	errs = r.db.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Username, user.Email, user.Password).Scan(&result.ID, &result.FirstName, &result.LastName, &result.Username, &result.Email, &result.CreatedAt, &result.UpdatedAt)
	if errs != nil {
		return errs, domain.User{}
	}
	return nil, result
}

func (r *userRepository) GetOne(user domain.User, option string) (result domain.User, err error) {
	sqlStatement := "SELECT id, first_name, last_name, username, email, created_at, updated_at FROM users WHERE"
	var row *sql.Row

	if option == "email" {
		sqlStatement += " email = $1"
		row = r.db.QueryRow(sqlStatement, user.Email)
	}

	if option == "username" {
		sqlStatement += " username = $1"
		row = r.db.QueryRow(sqlStatement, user.Username)
	}

	switch err := row.Scan(&result.ID, &result.FirstName, &result.LastName, &result.Username, &result.Email, &result.CreatedAt, &result.UpdatedAt); err {
	case sql.ErrNoRows:
		return domain.User{}, errors.New(fmt.Sprintf("record not found"))
	case nil:
		return result, nil
	default:
		return domain.User{}, err
	}
}

func (r *userRepository) GetOneForAuth(user domain.User) (result domain.User, err error) {
	sqlStatement := "SELECT id, first_name, last_name, username, email, password, created_at, updated_at FROM users WHERE email = $1"

	row := r.db.QueryRow(sqlStatement, user.Email)
	switch err := row.Scan(&result.ID, &result.FirstName, &result.LastName, &result.Username, &result.Email, &result.Password, &result.CreatedAt, &result.UpdatedAt); err {
	case sql.ErrNoRows:
		return domain.User{}, errors.New(fmt.Sprintf("record not found"))
	case nil:
		return result, nil
	default:
		return domain.User{}, err
	}
}

func (r *userRepository) UpdatePassword(user domain.User) error {
	sqlStatement := "UPDATE users SET password=$1 WHERE email = $2"
	user.HashPassword()
	err := r.db.QueryRow(sqlStatement, user.Password, user.Email)
	return err.Err()
}
