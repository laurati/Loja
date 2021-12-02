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

func InsertProducts(c *gin.Context) {
	var novoProduto models.Produto

	if err := c.BindJSON(&novoProduto); err != nil {
		return
	}

	models.CriaNovoProduto(novoProduto.Nome, novoProduto.Preco, int(novoProduto.Quantidade))

	c.IndentedJSON(http.StatusCreated, novoProduto)
}
