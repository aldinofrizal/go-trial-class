package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type EmailRequest struct {
	Email       string `json:"email"`
	Address     string `json:"address"`
	ProductName string `json:"product_name"`
}

func MailSuccessCreateOrder(ctx *gin.Context) {
	var emailRequest EmailRequest

	if err := ctx.ShouldBindJSON(&emailRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "trial-class@mail.com")
	m.SetHeader("To", emailRequest.Email)

	m.SetHeader("Subject", "Trial Class Order")
	m.SetBody("text/html", fmt.Sprintf("terima kasih telah melakukan order pada mini ecommerce trial mini class, product dengan nama %s akan dikirimkan ke alamat %s secepatnya", emailRequest.ProductName, emailRequest.Address))

	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")

	if err := d.DialAndSend(m); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "email berhasil terkirim",
	})
}
