package entity

import (
	"fmt"
	"time"
)

type Order struct {
	ID           uint
	ProductId    int
	BuyerEmail   string
	BuyerAddress string
	OrderDate    time.Time
	Product      Product
}

func (o *Order) PrintDetail() {
	fmt.Println("Email pembeli: ", o.BuyerEmail)
	fmt.Println("Alamat pembeli: ", o.BuyerAddress)
	fmt.Println("Produk: ", o.Product.Name, "Tanggal order", o.OrderDate)
	fmt.Println("")
}
