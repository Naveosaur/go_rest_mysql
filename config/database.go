package config

import (
	"fmt"
	"go-rest-sql/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	connectionStr := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_URL, ENV.DB_DATABASE, "charset=utf8mb4&parseTime=true&loc=Asia%2FJakarta")

	db, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})

	if err != nil {
		panic("Database connection failed")	
	}

	
	db.AutoMigrate(&models.Author{}, &models.Book{})

	DB = db
	log.Println("Database connected")

}