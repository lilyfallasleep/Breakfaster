package mybot

import (
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

// ReplyFlex is a method for sending back flex message
func (app *BreakFaster) replyFlex(replyToken string, title string, flexMsgFunc func() linebot.FlexContainer, disableCache bool) error {
	var flexMsg linebot.FlexContainer
	if disableCache {
		flexMsg = flexMsgFunc()
	} else {
		flexMsg = app.cacheWrapper(title, flexMsgFunc)
	}

	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewFlexMessage(title, flexMsg),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *BreakFaster) replyText(replyToken, text string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *BreakFaster) replyCancelConfirmBox(replyToken string) error {
	cancelParams := "cancel_order"
	template := linebot.NewConfirmTemplate(
		"是否確定取消訂單?",
		linebot.NewPostbackAction("Yes", cancelParams, "", ""),
		linebot.NewPostbackAction("No", "dummy", "", ""),
	)
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTemplateMessage("取消訂單確認", template),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *BreakFaster) replyOrderConfirmCard(replyToken, lineUID string, start, end time.Time) error {
	confirmCard, err := app.NewConfirmCard(lineUID, start, end)
	if err != nil {
		return err
	}
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewFlexMessage("點餐紀錄", confirmCard),
	).Do(); err != nil {
		return err
	}
	return nil
}

// BroadcastFlex is a method for broacasting flex message to all friends
func (app *BreakFaster) BroadcastFlex(title string, flexMsg linebot.FlexContainer) error {
	if _, err := app.bot.BroadcastMessage(
		linebot.NewFlexMessage(title, flexMsg)).Do(); err != nil {
		return err
	}
	return nil
}

// SendDirectFlex is a method for sending flex message directly to an user
func (app *BreakFaster) SendDirectFlex(lineUID, title string, flexMsg linebot.FlexContainer) error {
	if _, err := app.bot.PushMessage(lineUID, linebot.NewFlexMessage(title, flexMsg)).Do(); err != nil {
		return err
	}
	return nil
}
