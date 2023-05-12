package cli

import (
	"fmt"
	"os"
)

func ErrorHandler(msg string) {
	fmt.Println("Terjadi kesalahan dalam aplikasi")
	fmt.Println(msg)

	var input string
	fmt.Println("Tekan (m) untuk kembali ke halaman utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err.Error())
	}

	switch input {
	case "m":
		MainMenu()
	case "q":
		fmt.Println("Terima kasih telah menggunakan aplikasi Mini Ecommerce")
		os.Exit(1)
	default:
		ErrorHandler(msg)
	}
}
