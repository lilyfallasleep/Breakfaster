package mybot

import (
	c "breakfaster/config"

	"github.com/line/line-bot-sdk-go/linebot"
)

// OrderPageURI is the liff page URI
var OrderPageURI string

// BreakFastPusher is the type for message pushing
type BreakFastPusher struct {
	bot *linebot.Client
}

// BreakFaster is the type for bot
type BreakFaster struct {
	bot *linebot.Client
	svc *WebhookService
}

// NewLineBot is the factory for bot instance
func NewLineBot(config *c.Config) (*linebot.Client, error) {
	// initialize liff page URI
	OrderPageURI = config.OrderPageURI

	bot, err := linebot.New(
		config.ChannelSecret,
		config.AccessToken,
	)
	if err != nil {
		return nil, err
	}
	return bot, nil
}

// NewBreakFastPusher is a factory for message pusher instance
func NewBreakFastPusher(bot *linebot.Client) (BreakFastPushBot, error) {
	return &BreakFastPusher{
		bot: bot,
	}, nil
}

// NewBreakFaster is a factory for webhook instance
func NewBreakFaster(bot *linebot.Client, svc *WebhookService) (BreakFastBot, error) {
	return &BreakFaster{
		bot: bot,
		svc: svc,
	}, nil
}
