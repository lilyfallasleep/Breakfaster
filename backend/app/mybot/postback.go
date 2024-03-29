package mybot

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func (app *BreakFaster) nextStep(replyToken string, source *linebot.EventSource, data string) error {
	switch data {
	case "check_order":
		start, end := app.svc.timer.GetNextWeekInterval()
		confirmCard, err := app.NewConfirmCard(source.UserID, start, end)
		if err != nil {
			return err
		}
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewFlexMessage("點餐紀錄", confirmCard),
		).Do(); err != nil {
			return err
		}
	case "cancel_order":
		start, end := app.svc.timer.GetNextWeekInterval()
		if err := app.svc.orderRepo.DeleteOrdersByLineUID(source.UserID, start, end); err != nil {
			return err
		}
		if err := app.replyText(replyToken, "訂單已取消😎"); err != nil {
			return err
		}
	case "dummy":
		if err := app.replyText(replyToken, "請繼續使用 Breakfaster！"); err != nil {
			return err
		}
	case "rule":
		if err := app.replyFlex(replyToken, "點餐規則", NewWelcomeCard, false); err != nil {
			return err
		}
	default:
		app.svc.logger.Errorf("Unknown action: %v", data)
	}
	return nil
}
