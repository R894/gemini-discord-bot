package main

import (
	"log"

	"github.com/R894/gemini-test/gemini"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	message := "Write me a hello world program in Rust"
	part, err := gemini.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(part)
}
