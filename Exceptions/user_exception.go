package apperror

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email já cadastrado")
	ErrInvalidPassword    = errors.New("senha invalida")
	ErrInvalidCredentials = errors.New("mail ou senha inválidos")
)
