package controller

import (
	"net/http"
	"productcrud/usecase"

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
	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
