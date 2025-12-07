package controller

import (
	"net/http"
	"productcrud/model"
	"productcrud/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	// Aqui sera colocado o useCase
	productUsecase usecase.ProductUsecase
}

// Aqui seria como se fosse uma classe;
func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

// Aqui seria como se fosse um metodo de uma classe;
func (p *productController) GetProducts(ctx *gin.Context) {
	userIDAny, exists := ctx.Get("userId")
	if !exists {
		ctx.JSON(401, gin.H{"error": "Usuario nao autenticado!!"})
		return
	}

	// Pegando o id do usuario autenticado
	userID := userIDAny.(int)
	_ = userID

	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("pdId")
	if id == "" {
		response := model.Responde{Message: "Id não pode ser nulo!"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Tentando converter para numero
	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Responde{Message: "Id informado invalido!"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := p.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if product == nil {
		response := model.Responde{Message: "Produto nao encontrado!"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("pdId")
	if id == "" {
		response := model.Responde{Message: "Id não pode ser nulo!"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Tentando converter para numero
	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Responde{Message: "Id informado invalido!"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product

	err = ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	product.ID = productId

	updatedPd, err := p.productUsecase.UpdateProduct(&product)
	if err != nil {
		println("Erro aqui")
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if updatedPd == nil {
		response := model.Responde{Message: "Produto nao encontrado!"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)

}
