package repository

import (
	"database/sql"
	"errors"
	"fmt"
	apperror "productcrud/Exceptions"
	"productcrud/model"

	"github.com/lib/pq"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) RegisterUser(user *model.User) (int, error) {
	query := "INSERT INTO users(name, email, password) VALUES ($1, $2, $3) RETURNING id"

	var id int

	err := ur.connection.QueryRow(query, user.Name, user.Email, user.Password).Scan(&id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" && pqErr.Constraint == "users_email_key" {
				fmt.Println(apperror.ErrEmailAlreadyExists)
				return 0, apperror.ErrEmailAlreadyExists
			}
		}

		return 0, err
	}

	return id, nil
}

func (ur *UserRepository) LoginUser(user *model.User) (*model.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email=$1"

	var userLogin model.User

	err := ur.connection.QueryRow(query, user.Email).Scan(&userLogin.ID, &userLogin.Name, &userLogin.Email, &userLogin.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrInvalidCredentials
		}
		fmt.Println(err)
		return nil, err
	}

	return &userLogin, nil

}
