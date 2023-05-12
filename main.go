package main

import (
	"mini-ecommerce/cli"
	"mini-ecommerce/config"
)

func main() {
	config.DBConnect()
	cli.MainMenu()
}
