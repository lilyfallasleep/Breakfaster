package mybot

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Callback is a callback method for handling http requests
func (app *BreakFaster) Callback(w http.ResponseWriter, r *http.Request) {
	events, err := app.bot.ParseRequest(r)
	if err != nil {
		app.svc.logger.Error(err)
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
					if err := app.replyFlex(event.ReplyToken, "操作選單", NewMenu, false); err != nil {
						app.svc.logger.Error(err)
					}
				case "點餐紀錄":
					start, end := app.svc.timer.GetNextWeekInterval()
					if err := app.replyOrderConfirmCard(event.ReplyToken, event.Source.UserID, start, end); err != nil {
						app.svc.logger.Error(err)
					}
				case "取消訂單":
					if err := app.replyCancelConfirmBox(event.ReplyToken); err != nil {
						app.svc.logger.Error(err)
					}
				default:
					if err := app.Predict(event.ReplyToken, event.Source.UserID, message.Text); err != nil {
						app.svc.logger.Error(err)
					}
				}
			default:
				app.svc.logger.Errorf("Unknown message: %v", message)
			}
		case linebot.EventTypeFollow:
			if err := app.replyFlex(event.ReplyToken, "點餐規則", NewWelcomeCard, false); err != nil {
				app.svc.logger.Error(err)
			}
		case linebot.EventTypePostback:
			data := event.Postback.Data

			if err := app.nextStep(event.ReplyToken, event.Source, data); err != nil {
				app.svc.logger.Error(err)
			}
		default:
			app.svc.logger.Errorf("Unknown event: %v", event)
		}
	}
}
