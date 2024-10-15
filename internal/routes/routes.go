package routes

import (
	"ProductService/internal/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(gin *gin.Engine) {
	gin.GET("/test", controllers.Test)
	gin.POST("/createProduct", controllers.CreateProduct)
	gin.GET("/getProduct/:id", controllers.GetProduct)
	gin.GET("/getAllProducts", controllers.GetAllProducts)
	gin.PUT("/updateProduct/:id", controllers.UpdateProduct)
	gin.DELETE("/deleteProduct/:id", controllers.DeleteProduct)
}
