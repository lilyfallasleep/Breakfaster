package mybot

import (
	"net/http"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

// BreakFastBot is the Breakfaster interface
type BreakFastBot interface {
	Callback(w http.ResponseWriter, r *http.Request)
	Predict(replyToken, lineUID, text string) error
	NewConfirmCard(lineUID string, start, end time.Time) (linebot.FlexContainer, error)
}

// BreakFastPushBot is the BreakFastPusher interface
type BreakFastPushBot interface {
	BroadcastFlex(title string, flexMsg linebot.FlexContainer) error
	SendDirectFlex(lineUID, title string, flexMsg linebot.FlexContainer) error
}
