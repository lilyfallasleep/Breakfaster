package mybot

import (
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

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
