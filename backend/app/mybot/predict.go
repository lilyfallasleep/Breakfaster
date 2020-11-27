package mybot

// Predict returns message according to the given text
func (app *BreakFaster) Predict(replyToken, lineUID, text string) error {
	prediction, err := app.svc.ar.Predict(text)
	if err != nil {
		return err
	}
	switch prediction {
	case "問題回報":
		resp := "請點擊以下連結回報問題：\n\n" + OrderPageURI + "/report"
		if err := app.replyText(replyToken, resp); err != nil {
			return err
		}
	case "取消訂單":
		if err := app.replyCancelConfirmBox(replyToken); err != nil {
			return err
		}
	case "點餐紀錄":
		start, end := app.svc.timer.GetNextWeekInterval()
		if err := app.replyOrderConfirmCard(replyToken, lineUID, start, end); err != nil {
			return err
		}
	case "規則":
		if err := app.replyFlex(replyToken, "點餐規則", NewWelcomeCard, false); err != nil {
			return err
		}
	default:
		resp := "請點擊以下連結開始點餐！\n\n" + OrderPageURI
		if err := app.replyText(replyToken, resp); err != nil {
			return err
		}
	}
	return nil
}
