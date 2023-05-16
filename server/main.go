package main

import (
	"trial-class-api/config"
	"trial-class-api/controller"

	_ "trial-class-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Trial Class Mini Ecommerce
// @version         1.0
// @description     Dokomentasi REST API project Mini Ecommerce II

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:8000
func main() {
	s := gin.Default()
	config.DBConnect()

	s.GET("/products", controller.GetProductHandler)
	s.POST("/orders", controller.PostOrderHandler)
	s.GET("/orders", controller.GetOrderHandler)
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s.Run(":8000")
}
