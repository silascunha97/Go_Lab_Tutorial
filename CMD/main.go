package main

import (
	"localhost/Go-Lab-API/controller"
	"localhost/Go-Lab-API/db"
	"localhost/Go-Lab-API/repository"
	"localhost/Go-Lab-API/usecase"

	"github.com/gin-gonic/gin" // Importa o pacote gin, que é um framework web para Go
)

type productController struct {
	productUsecase usecase.ProductUseCase
}

func main() {
	// Função principal onde a execução do programa começa

	server := gin.Default()
	// Cria uma nova instância do servidor Gin com as configurações padrão

	dbConnection, err := db.ConnectDB() // Conecta ao banco de dados
	if err != nil {
		// Se ocorrer um erro ao conectar ao banco de dados
		panic(err)
		// Encerra o programa e exibe o erro

	}

	//camada usecase
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Cria uma nova instância de ProductRepository passando a conexão com o banco de dados como parâmetro
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)

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
