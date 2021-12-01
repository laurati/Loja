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
