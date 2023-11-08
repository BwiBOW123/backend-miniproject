package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}

}
