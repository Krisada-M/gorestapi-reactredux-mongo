package config

import (
	"log"

	"github.com/joho/godotenv"
)

func Envload() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env not found")
	}
}
