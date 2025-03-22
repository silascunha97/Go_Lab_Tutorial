package controller

import (
	"localhost/Go-Lab-API/usecase"
	"net/http" // Importa o pacote net/http para lidar com códigos de status HTTP

	"github.com/gin-gonic/gin" // Importa o pacote gin para criar manipuladores de rotas
	// Importa o pacote model do projeto para acessar o modelo Product
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

	products, err := p.productUseCase.GetProducts() // Chama o método GetProducts do productUseCase para obter a lista de produtos
	if err != nil {                                 // Verifica se ocorreu um erro ao obter a lista de produtos
		ctx.JSON(http.StatusInternalServerError, err) // Responde à requisição com um status HTTP 500 (Internal Server Error) e uma mensagem de erro em formato JSON
	}

	// Responde à requisição com um status HTTP 200 (OK) e a lista de produtos em formato JSON
	ctx.JSON(http.StatusOK, products)
}
