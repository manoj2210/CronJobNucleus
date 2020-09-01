package main

import (
	"Webhooks/internal/app"
	"Webhooks/internal/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	githubClient:= config.NewGithubClient()
	app.StartApplication(githubClient)
}
