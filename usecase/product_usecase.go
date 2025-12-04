package usecase

import (
	"productcrud/model"
	"productcrud/repository"
)

// Para que seja vidivel em outro pacotes ela deve estar com a primeira letra maiuscula;
type ProductUsecase struct {
	// repository
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
