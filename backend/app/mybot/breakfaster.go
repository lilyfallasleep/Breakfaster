package mybot

import (
	"breakfaster/mybot/autoreply"
	"breakfaster/pkg/ordertime"
	"breakfaster/repository/dao"

	"github.com/line/line-bot-sdk-go/linebot"

	c "breakfaster/config"
	"breakfaster/infrastructure/cache"
)

// BreakFaster is the type for bot
type BreakFaster struct {
	bot          *linebot.Client
	ar           *autoreply.AutoReplier
	msgCache     cache.GeneralCache
	orderRepo    *dao.OrderRepository
	timer        *ordertime.OrderTimer
	OrderPageURI string
}

// NewBreakFaster is a factory for bot instance
func NewBreakFaster(config *c.Config, ar *autoreply.AutoReplier, msgCache cache.GeneralCache, orderRepo *dao.OrderRepository, timer *ordertime.OrderTimer) (*BreakFaster, error) {
	bot, err := linebot.New(
		config.ChannelSecret,
		config.AccessToken,
	)
	if err != nil {
		return nil, err
	}

	return &BreakFaster{
		bot:          bot,
		ar:           ar,
		msgCache:     msgCache,
		orderRepo:    orderRepo,
		timer:        timer,
		OrderPageURI: config.OrderPageURI,
	}, nil
}
