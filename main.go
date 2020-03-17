package main

import (
	"./fetcher"
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(errors.New("err not nil"))
	}
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fetcher.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))

	fetcher.SetConfig(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))

	fetcher.Run()

	err = fetcher.Close()
	if err != nil {
		panic(err)
	}
}
