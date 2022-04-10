package store

import (
	"auth/internal/model"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
)

// UserRepository
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) (*model.User, error) {

	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	q := `
	INSERT INTO users
	(email, encrypted_password)
	VALUES 
	($1, $2)
	returning id
	
`

	if err := r.store.Pool.QueryRow(context.Background(), q, u.Email, u.EncryptedPassword).
		Scan(
			&u.ID,
		); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf(
				"SQL Error: %s, Detail: %s, Where: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.SQLState()))
			return nil, newErr
		} else {
			return nil, err
		}
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {

	u := &model.User{}

	q := `
	SELECT id, email, encrypted_password
	FROM users
	WHERE email = $1
`
	if err := r.store.Pool.QueryRow(context.Background(), q, email).
		Scan(
			&u.ID,
			&u.Email,
			&u.EncryptedPassword); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf(fmt.Sprintf(
				"SQL Error: %s, Detail: %s, Where: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.SQLState()))
			return nil, newErr
		} else {
			return nil, err
		}
	}

	return u, nil
}
