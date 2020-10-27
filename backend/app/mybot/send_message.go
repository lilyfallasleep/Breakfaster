package mybot

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// BroadcastFlex is a method for broacasting flex message to all friends
func (pusher *BreakFastPusher) BroadcastFlex(title string, flexMsg linebot.FlexContainer) error {
	if _, err := pusher.bot.BroadcastMessage(
		linebot.NewFlexMessage(title, flexMsg)).Do(); err != nil {
		return err
	}
	return nil
}

// SendDirectFlex is a method for sending flex message directly to an user
func (pusher *BreakFastPusher) SendDirectFlex(lineUID, title string, flexMsg linebot.FlexContainer) error {
	if _, err := pusher.bot.PushMessage(lineUID, linebot.NewFlexMessage(title, flexMsg)).Do(); err != nil {
		return err
	}
	return nil
}
