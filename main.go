package main

import (
	"./dadtaoke"
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

	dadtaoke.Connect(os.Getenv("DB_DRIVER"), os.Getenv("DB_DSN"))

	dadtaoke.SetConfig(os.Getenv("APP_KEY"), os.Getenv("APP_SECRET"))

	dadtaoke.Run()

	err = dadtaoke.Close()
	if err != nil {
		panic(err)
	}
}
