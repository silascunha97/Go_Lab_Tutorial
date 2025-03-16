package controller

import (
	"net/http" // Importa o pacote net/http para lidar com códigos de status HTTP

	"github.com/gin-gonic/gin"  // Importa o pacote gin para criar manipuladores de rotas
	"localhost.com/GoLab/model" // Importa o pacote model do projeto para acessar o modelo Product
	"localhost.com/GoLab/usecase"
)

// Define a estrutura productController, que será usada para agrupar os métodos do controlador de produtos
type productController struct {
	productUseCase usecase.ProductUseCase
}

// NewProductController cria uma nova instância de productController
func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
	// Retorna uma nova instância de productController
}

// GetProducts é um método do productController que lida com a rota GET para obter produtos
func (p *productController) GetProducts(ctx *gin.Context) {
	// Cria uma lista de produtos com um produto de exemplo
	products := []model.Product{
		{
			ID:    "3",            // ID do produto
			Name:  "Batata Frita", // Nome do produto
			Price: 19.99,          // Preço do produto
		},
	}

	// Responde à requisição com um status HTTP 200 (OK) e a lista de produtos em formato JSON
	ctx.JSON(http.StatusOK, products)
}
