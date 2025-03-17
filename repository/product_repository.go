package repository

import (
	"fmt" // Importa o pacote fmt para formatação de strings e impressão

	"database/sql" // Importa o pacote sql para trabalhar com bancos de dados SQL

	"localhost.com/GoLab/model" // Importa o pacote model que contém a definição do modelo Product
)

// Define a estrutura ProductRepository que contém uma conexão com o banco de dados
type ProductRepository struct {
	connection *sql.DB // Campo connection que é um ponteiro para uma instância de sql.DB
}

// Função que cria uma nova instância de ProductRepository com a conexão fornecida
func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection, // Inicializa o campo connection com a conexão fornecida
	}
}

// Método GetProducts que retorna uma lista de produtos e um erro, se houver
func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	// Define a consulta SQL para selecionar id, nome do produto e preço da tabela product
	query := "SELECT id, product_name, price FROM product"
	// Executa a consulta no banco de dados
	rows, err := pr.connection.Query(query)
	if err != nil {
		// Se houver um erro na execução da consulta, imprime o erro e retorna uma lista vazia e o erro
		fmt.Printf(err.Error())
		return []model.Product{}, err
	}

	// Declara uma lista de produtos para armazenar os resultados da consulta
	var productList []model.Product
	// Declara uma variável productObj do tipo model.Product para armazenar cada linha da consulta
	var productObj model.Product

	// Itera sobre as linhas retornadas pela consulta
	for rows.Next() {
		// Escaneia os valores da linha atual para os campos de productObj
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			// Se houver um erro ao escanear a linha, imprime o erro e retorna uma lista vazia e o erro
			fmt.Printf(err.Error())
			return []model.Product{}, err
		}

		// Adiciona o productObj à lista de produtos
		productList = append(productList, productObj)
	}

	// Fecha o conjunto de resultados
	rows.Close()

	// Retorna a lista de produtos e nil indicando que não houve erro
	return productList, nil
}
