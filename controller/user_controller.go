package controller

import (
	"errors"
	apperror "productcrud/Exceptions"
	"productcrud/auth"
	"productcrud/model"
	"productcrud/usecase"
	"regexp"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

// Cria a "classe do controller" é o construtor
func NewUserController(usecase usecase.UserUsecase) UserController {
	return UserController{
		userUsecase: usecase,
	}
}

// ctx é um objeto único que tem tudo: request + response + helpers
func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	var body model.RegisterUserDTO
	if err := ctx.BindJSON(&body); err != nil {
		//ctx.JSON envia a resposta já define status, headers e body.
		// o gin.H é basicamente: map[string]interface{}{"error": "FORMATO JSON INVALIDO!",}, é um map em go, dizendo que a chave é string e o valor é qualquer um

		ctx.JSON(400, gin.H{"error": "FORMATO JSON INVALIDO!"})
		return
	}

	if len(body.Name) <= 3 {
		ctx.JSON(400, gin.H{"error": "Nome deve ter masi que 3 caracteres"})
		return
	}

	if len(body.Password) <= 8 {
		ctx.JSON(400, gin.H{"error": "Sennha muito curta!"})
		return
	}

	if !(emailRegex.MatchString(body.Email)) {
		ctx.JSON(400, gin.H{"error": "Email Invalido!"})
		return
	}

	// Retorna apeas o id do usuario criado!!
	id, err := uc.userUsecase.RegisterUser(&body)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"sucess": true, "id": id})

}

func (uc *UserController) LoginUser(ctx *gin.Context) {
	var body model.RegisterUserDTO
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": "FORMATO JSON INVALIDO!"})
		return
	}
	user, err := uc.userUsecase.LoginUser(&body)
	if err != nil {
		if errors.Is(err, apperror.ErrInvalidPassword) {
			ctx.JSON(400, gin.H{"error": "Senha inválida"})
			return
		}

		if errors.Is(err, apperror.ErrInvalidCredentials) {
			ctx.JSON(400, gin.H{"error": "Email ou senha inválidos"})
			return
		}

		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Gerar o token com JWT
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(200, gin.H{"sucess": false, "error": "Erro ao gerar o token!!"})
		return
	}

	//retornando o token para o front salvar e mandar nas requisicoes
	ctx.JSON(200, gin.H{
		"sucess":  true,
		"message": "Login feito com sucesso!",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})

}
