package usecase

import (
	apperror "productcrud/Exceptions"
	"productcrud/model"
	"productcrud/repository"
	"strings"
)

type UserUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return UserUsecase{
		repository: ur,
	}
}

func (uu *UserUsecase) RegisterUser(input *model.RegisterUserDTO) (int, error) {
	//Limpeza, hash, validacao, normalizacao...
	email := strings.TrimSpace(strings.ToLower(input.Email))
	name := strings.TrimSpace(input.Name)

	//Faria o Hash
	//Hashed password

	// Sem & → você cria um valor
	// Com & → você cria um ponteiro, que guarda o endereço onde o valor está
	user := &model.User{
		Name:     name,
		Email:    email,
		Password: input.Password, // Usando assim por enquanto
	}

	id, err := uu.repository.RegisterUser(user)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (uu *UserUsecase) LoginUser(input *model.RegisterUserDTO) (*model.User, error) {

	//Faria o hash da senha:
	// Password

	userInput := &model.User{
		Email:    input.Email,
		Password: input.Password,
	}

	user, err := uu.repository.LoginUser(userInput)
	if err != nil {
		return nil, err
	}

	//Verificação de senha do banco com o input
	if userInput.Password != user.Password {
		return nil, apperror.ErrInvalidPassword
	}

	return user, nil

}

// Controller → manda dados simples → UseCase

// UseCase → cria entidade → passa ponteiro para Repository

// Repository → preenche ID, created_at, etc → retorna ponteiro
