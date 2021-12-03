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
	router.DELETE("/delete/:id", controllers.Delete)
	router.PUT("/edit/:id", controllers.Edit)

	router.Run("localhost:8080")
}
