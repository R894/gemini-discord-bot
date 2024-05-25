package main

import (
	"log"

	"github.com/R894/gemini-test/discord"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	println("Starting...")
	discord.Start()
}
