package entity

import "time"

type Order struct {
	ID           uint      `json:"id"`
	ProductId    int       `json:"product_id" binding:"required"`
	BuyerEmail   string    `json:"buyer_email" binding:"required,email"`
	BuyerAddress string    `json:"buyer_address" binding:"required"`
	OrderDate    time.Time `json:"order_date"`
	Product      Product   `json:"product"`
}

type OrderRequest struct {
	ProductId    int    `json:"product_id" binding:"required"`
	BuyerEmail   string `json:"buyer_email" binding:"required,email"`
	BuyerAddress string `json:"buyer_address" binding:"required"`
}
