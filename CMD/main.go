package main

import (
	"github.com/gin-gonic/gin" // Importa o pacote gin, que é um framework web para Go
	"localhost.com/GoLab/controller"
	"localhost.com/GoLab/usecase"
)

type productController struct {
	productUsecase usecase.ProductUseCase
}

func main() {
	// Função principal onde a execução do programa começa

	server := gin.Default()
	// Cria uma nova instância do servidor Gin com as configurações padrão

	ProductUseCase := usecase.NewProductUsecase()
	// Cria uma nova instância de ProductUsecase

	ProductController := controller.NewProductController(ProductUseCase)
	// Cria uma nova instância de ProductController passando ProductUsecase como parâmetro

	server.GET("/ping", func(ctx *gin.Context) {
		// Define uma rota GET para o caminho "/ping"
		// Quando essa rota é acessada, a função anônima é executada

		ctx.JSON(200, gin.H{
			"message": "pong",
		})
		// Responde à requisição com um status HTTP 200 (OK) e um JSON contendo {"message": "pong"}
	})

	server.GET("/products", ProductController.GetProducts)
	/*
	   Define uma rota GET para o caminho "/products"
	   Quando essa rota é acessada, o método GetProducts do controlador de produtos é executado
	   O método GetProducts retorna uma lista de produtos em formato JSON
	*/

	server.Run(":8080")
	// Inicia o servidor na porta 8080
}
