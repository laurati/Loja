package controllers

import (
	"net/http"
	"produtos-gin/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	todosOsProdutos := models.BuscaTodosOsProdutos()

	c.IndentedJSON(http.StatusOK, todosOsProdutos)

}

//outro jeito defazer usando a função BuscaTodosOsProdutos()
// func GetProductByID(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))

// 	for _, a := range models.BuscaTodosOsProdutos() {
// 		if a.Id == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
// }

func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	prod := models.BuscaProdutoPorId(id)

	c.IndentedJSON(http.StatusOK, prod)
}

func InsertProducts(c *gin.Context) {
	var novoProduto models.Produto

	if err := c.BindJSON(&novoProduto); err != nil {
		return
	}

	models.CriaNovoProduto(novoProduto.Nome, novoProduto.Preco, int(novoProduto.Quantidade))

	c.IndentedJSON(http.StatusCreated, novoProduto)
}
