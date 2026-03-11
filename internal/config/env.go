package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(){
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: .env load fail(%v)", err)
	}
}