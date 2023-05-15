package controller

import (
	"net/http"
	"time"

	"trial-class-api/config"
	"trial-class-api/entity"
	"trial-class-api/helpers"

	"github.com/gin-gonic/gin"
)

func GetProductHandler(ctx *gin.Context) {
	var products []entity.Product

	if err := config.DB.Find(&products).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "OK",
		"products": products,
	})
}

func PostOrderHandler(ctx *gin.Context) {
	var order entity.Order

	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var product entity.Product
	err := config.DB.Where("id = ?", order.ProductId).First(&product).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	order.OrderDate = time.Now()
	if err := config.DB.Create(&order).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	go helpers.SendMail(order.BuyerEmail, order.BuyerAddress, product.Name)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success create product",
	})
}

func GetOrderHandler(ctx *gin.Context) {
	var orders []entity.Order

	if err := config.DB.Preload("Product").Find(&orders).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"orders":  orders,
	})
}
