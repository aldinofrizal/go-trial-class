package main

import (
	"trial-class-mailer/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()

	s.POST("/mail", controller.MailSuccessCreateOrder)

	s.Run(":8001")
}
