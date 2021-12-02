package routes

import (
	"produtos-gin/controllers"

	"github.com/gin-gonic/gin"
)

func CarregaRotas() {

	router := gin.Default()
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProductByID)
	router.POST("/products", controllers.InsertProducts)

	router.Run("localhost:8080")
}
