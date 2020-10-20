package mybot

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Callback is a callback method for handling http requests
func (app *BreakFaster) Callback(w http.ResponseWriter, r *http.Request) {
	events, err := app.bot.ParseRequest(r)
	if err != nil {
		log.Print(err)
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
				case "點餐":
					if err := app.replyFlex(event.ReplyToken, "操作選單", app.NewMenu, false); err != nil {
						log.Print(err)
					}
				case "點餐紀錄":
					start, end := app.timer.GetNextWeekInterval()
					if err := app.replyOrderConfirmCard(event.ReplyToken, event.Source.UserID, start, end); err != nil {
						log.Print(err)
					}
				case "取消訂單":
					if err := app.replyCancelConfirmBox(event.ReplyToken); err != nil {
						log.Print(err)
					}
				default:
					if err := app.Predict(event.ReplyToken, event.Source.UserID, message.Text); err != nil {
						log.Print(err)
					}
				}
			default:
				log.Printf("Unknown message: %v", message)
			}
		case linebot.EventTypeFollow:
			if err := app.replyFlex(event.ReplyToken, "點餐規則", app.NewWelcomeCard, false); err != nil {
				log.Print(err)
			}
		case linebot.EventTypePostback:
			data := event.Postback.Data

			if err := app.nextStep(event.ReplyToken, event.Source, data); err != nil {
				log.Print(err)
			}
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}
