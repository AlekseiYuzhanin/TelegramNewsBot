package main

import (
	"app/internal/config"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
)

func main(){
	botAPI, err := tgbotapi.NewBotAPI(config.Get().TelegramBotToken)
	if err != nil{
		log.Printf("failed to create bot : %v", err)
		return
	}

	db, err := sqlx.Connect("postgres", config.Get().DatabaseDSN)
	if err != nil{
		log.Printf("failed to connect db : %v", err)
		return
	}

	defer db.Close()

	var (
		
	)
}