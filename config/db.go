package config

import (
	"fmt"
	"log"
	"mini-ecommerce/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	dsn := "host=localhost user=postgres password=123456 dbname=trial_class_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err.Error())
	} else {
		fmt.Println("connected to db")
		DB = db
	}

	db.AutoMigrate(&entity.Product{}, &entity.Order{})
}
