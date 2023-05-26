package cli

import (
	"bufio"
	"fmt"
	"mini-ecommerce/config"
	"mini-ecommerce/entity"
	"mini-ecommerce/helpers"
	"os"
)

func ListOrder() {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)
	var orders []entity.Order

	err := config.DB.Preload("Product").Find(&orders).Error
	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("List order ---")
	for _, order := range orders {
		order.PrintDetail()
	}

	var input string
	fmt.Println("Tekan (key apapun) untuk kembali ke halaman utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	// _, err = fmt.Scanln(&input)
	input, _ = consoleReader.ReadString('\n')
	// if err != nil {
	// 	panic(err.Error())
	// }

	switch input {
	case "q":
		fmt.Println("Terima kasih telah menggunakan aplikasi Mini Ecommerce")
		os.Exit(1)
	default:
		MainMenu()
	}
}
