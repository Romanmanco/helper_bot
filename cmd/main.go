package main

import (
	"github.com/joho/godotenv"
	"helper_bot/internal/config"
	"helper_bot/internal/services"
	"log"
)

// init загрузка конфига перед стартом main()
func init() {
	log.Println("loading app env...")
	if err := godotenv.Load("app.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	log.Println("initializing app...")
	cfg := config.GetConfig()

	start := services.StartBotService{}

	err := start.RunBot(cfg.TelegramToken, cfg.WeatherToken)
	if err != nil {
		log.Fatal(err.Error())
	}
}
