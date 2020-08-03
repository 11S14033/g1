package main

import (
	"fmt"
	"g1/handlers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var app = handlers.App{}

func main() {
	Run()
}

func Run() {
	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error when read and get data from env, cause: %v", err)
	} else {
		fmt.Println("Getting data form env")
	}

	//initialize app
	app.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	app.Run(os.Getenv("API_PORT"))

}
