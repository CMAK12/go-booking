package main

import (
	"log"

	"go-booking/internal/app"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}
}

func main() {
	app.MustRun()
}
