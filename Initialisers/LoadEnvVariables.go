package initialisers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load("/home/ubuntu/CityViBe-ecommerce--Project/.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
