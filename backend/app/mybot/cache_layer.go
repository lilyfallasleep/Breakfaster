package mybot

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

func (app *BreakFaster) cacheWrapper(key string, fallbackFunc func() linebot.FlexContainer) linebot.FlexContainer {
	if flexMsg, found := app.svc.msgCache.Get(key); found {
		return flexMsg.(linebot.FlexContainer)
	}
	flexMsg := fallbackFunc()
	app.svc.msgCache.Set(key, flexMsg)
	return flexMsg
}
