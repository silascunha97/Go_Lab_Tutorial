package usecase

import (
	"localhost.com/GoLab/model"
	"localhost.com/GoLab/repository"
)

// Define a estrutura ProductUseCase, que será usada para agrupar os métodos do caso de uso de produtos
type ProductUseCase struct {
	repository repository.ProductRepository
}

// NewProductUsecase cria uma nova instância de ProductUseCase
func NewProductUsecase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
	// Retorna uma nova instância de ProductUseCase
}

// GetProducts é um método do ProductUseCase que retorna uma lista de produtos
func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
	// Retorna uma lista vazia de produtos e nenhum erro
}
