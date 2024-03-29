package config

import (
	"log"

	"github.com/spf13/viper"
)

// Seperti ENV
type Config struct {
	PORT string
	DB_USERNAME string
	DB_PASSWORD string
	DB_DATABASE string
	DB_URL string
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}

	log.Println("load successfully")
}