package controller

import (
	"net/http"
	"time"

	"trial-class-api/config"
	"trial-class-api/entity"
	"trial-class-api/helpers"

	"github.com/gin-gonic/gin"
)

// @Summary Get Product
// @Schemes Product
// @Description Get list of all available Products
// @Tags Product
// @Produce json
// @Success 200 {array} entity.Product
// @Router /products [get]
func GetProductHandler(ctx *gin.Context) {
	var products []entity.Product

	if err := config.DB.Find(&products).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// @Summary Post Order
// @Schemes Order
// @Description Create order to order specific product
// @Tags Order
// @Accept json
// @Produce json
// @Param data body entity.OrderRequest true "Order data"
// @Success 200 {string} success create order
// @Router /orders [post]
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
		"message": "success create order",
	})
}

// @Summary Get Order
// @Schemes Order
// @Description Get list of all orders
// @Tags Order
// @Produce json
// @Success 200 {array} entity.Order
// @Router /orders [get]
func GetOrderHandler(ctx *gin.Context) {
	var orders []entity.Order

	if err := config.DB.Preload("Product").Find(&orders).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
