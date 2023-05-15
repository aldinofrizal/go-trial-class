package main

import (
	"trial-class-api/config"
	"trial-class-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	config.DBConnect()

	s.GET("/products", controller.GetProductHandler)

	s.POST("/orders", controller.PostOrderHandler)
	s.GET("/orders", controller.GetOrderHandler)

	s.Run(":8000")
}
