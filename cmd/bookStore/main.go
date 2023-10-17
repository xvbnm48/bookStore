package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// TODO: MAKE INIT SERVER
func main() {
	fmt.Println("hello world")
}

type conf struct {
	DbHost string
	DbPort string
	DbName string
	DbUser string
	DbPwd  string
}

func parseLoadConfig() (*conf, error) {
	err := godotenv.Load("../../internal/config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := &conf{
		DbHost: os.Getenv("DB_HOST"),
		DbName: os.Getenv("DB_NAME"),
		DbPwd:  os.Getenv("DB_PASSWORD"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("DB_USER"),
	}

	return config, nil
}
