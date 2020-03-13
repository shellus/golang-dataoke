package main

import (
	"./fetcher"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fetcher.SetConfig(fetcher.Config{
		App_key:    os.Getenv("APP_KEY"),
		App_secret: os.Getenv("APP_SECRET"),
	})
	fetcher.Run()
}
