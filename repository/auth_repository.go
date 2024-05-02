package repository

import (
	"database/sql"

	"jdnielss.dev/cats-social-app/model"
	"jdnielss.dev/cats-social-app/model/dto"
)

type AuthRepository interface {
	GetUserData(payload dto.LoginRequestDTO) (model.User, error)
}

type authRepository struct {
	db *sql.DB
}

func (a *authRepository) GetUserData(payload dto.LoginRequestDTO) (model.User, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return model.User{}, err
	}

	var user model.User

	err = tx.QueryRow(`
    SELECT id, email, password FROM users 
    WHERE email=$1
   `, payload.Email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
	)
	return user, nil
}

func NewUserRepository(db *sql.DB) AuthRepository {
	return &authRepository{db: db}
}
