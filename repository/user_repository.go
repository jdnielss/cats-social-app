package repository

import (
	"database/sql"
	"errors"

	"jdnielss.dev/cats-social-app/model"
)

type UserRepository interface {
	Get(id string) (model.User, error)
	GetByEmail(email string) (bool, error)
	Create(payload model.User) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(`
		SELECT 
			id, first_name, last_name, email, username, role, photo, created_at, updated_at 
		FROM 
			users 
		WHERE 
			id = $1`, id).
		Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.Password,
			&user.CreatedAt,
		)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *userRepository) GetByEmail(email string) (bool, error) {
	var count int
	err := u.db.QueryRow(`SELECT COUNT(email) FROM users WHERE email = $1`, email).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (u *userRepository) Create(payload model.User) (model.User, error) {
	var user model.User

	// Check if email already exists
	emailExists, err := u.GetByEmail(payload.Email)
	if err != nil {
		return model.User{}, err // return error if there's an issue with the database
	}
	if emailExists {
		return model.User{}, errors.New("email already registered") // return error if email already exists
	}

	// Proceed with user creation if email doesn't exist
	// Your user creation logic here...
	user.Email = payload.Email
	user.Name = payload.Name
	user.Password = payload.Password
	// You should add more fields as per your model definition

	// Insert user into the database
	_, err = u.db.Exec("INSERT INTO users (email, name, password) VALUES ($1, $2, $3)", user.Email, user.Name, user.Password)
	if err != nil {
		return model.User{}, err // return error if there's an issue with inserting user into the database
	}

	// Return success response
	return user, err
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
