package cli

import (
	"bufio"
	"fmt"
	"mini-ecommerce/config"
	"mini-ecommerce/entity"
	"mini-ecommerce/helpers"
	"os"
	"time"
)

func ListProduct() {
	helpers.ClearScreen()

	consoleReader := bufio.NewReader(os.Stdin)
	var products []entity.Product
	err := config.DB.Find(&products).Error

	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("--- List order ---")
	for _, product := range products {
		product.PrintDetail()
	}

	var input string
	fmt.Println("Masukan Id product untuk melanjutkan order")
	fmt.Println("Tekan (m) untuk kembali ke halaman utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	input, _ = consoleReader.ReadString('\n')
	// _, err = fmt.Scanln(&input)
	// if err != nil {
	// 	panic(err.Error())
	// }

	switch input {
	case "m\n":
		MainMenu()
	case "q\n":
		fmt.Println("Terima kasih telah menggunakan aplikasi Mini Ecommerce")
		os.Exit(1)
	default:
		OrderProduct(input)
	}
}

func OrderProduct(id string) {
	helpers.ClearScreen()

	consoleReader := bufio.NewReader(os.Stdin)

	var product entity.Product
	err := config.DB.Where("ID = ?", id).First(&product).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	product.PrintDetail()

	var input string
	fmt.Println("Tekan (y) untuk melanjutkan order")
	fmt.Println("Tekan (m) untuk kembali ke halaman utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")
	input, _ = consoleReader.ReadString('\n')
	// _, err = fmt.Scanln(&input)
	// if err != nil {
	// 	panic(err.Error())
	// }

	switch input {
	case "y\n":
		CreateOrder(product)
	case "m\n":
		MainMenu()
	case "q\n":
		fmt.Println("Terima kasih telah menggunakan aplikasi Mini Ecommerce")
		os.Exit(1)
	default:
		OrderProduct(id)
	}
}

func CreateOrder(product entity.Product) {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)
	var email string
	var address string

	fmt.Println("untuk melakukan order silahkan melengkapi data berikut")
	fmt.Println("Masukan email: ")
	email, _ = consoleReader.ReadString('\n')
	// _, err := fmt.Scanln(&email)
	// if err != nil {
	// 	ErrorHandler(err.Error())
	// 	return
	// }

	fmt.Println("Masukan alamat: ")
	address, _ = consoleReader.ReadString('\n')
	// _, err = fmt.Scanln(&address)
	// if err != nil {
	// 	ErrorHandler(err.Error())
	// 	return
	// }

	order := entity.Order{
		ProductId:    int(product.ID),
		BuyerEmail:   email,
		BuyerAddress: address,
		OrderDate:    time.Now(),
	}

	err := config.DB.Create(&order).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("Order berhasil.")
	var input string
	fmt.Println("Tekan (key apapun) untuk kembali ke halaman utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err = fmt.Scanln(&input)
	if err != nil {
		MainMenu()
	}

	switch input {
	case "q\n":
		fmt.Println("Terima kasih telah menggunakan aplikasi Mini Ecommerce")
		os.Exit(1)
	default:
		MainMenu()
	}
}
