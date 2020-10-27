package mybot

import (
	"breakfaster/mybot/autoreply"
	"breakfaster/pkg/ordertime"
	"breakfaster/repository/dao"

	c "breakfaster/config"
	"breakfaster/infrastructure/cache"

	log "github.com/sirupsen/logrus"
)

// WebhookService is the proxy for funtionalities needed in webhook
type WebhookService struct {
	ar        *autoreply.AutoReplier
	msgCache  cache.GeneralCache
	orderRepo *dao.OrderRepository
	timer     *ordertime.OrderTimer
	logger    *log.Entry
}

// NewWebhookService is a factory for webhook service instance
func NewWebhookService(config *c.Config, ar *autoreply.AutoReplier, msgCache cache.GeneralCache, orderRepo *dao.OrderRepository, timer *ordertime.OrderTimer) (*WebhookService, error) {
	return &WebhookService{
		ar:        ar,
		msgCache:  msgCache,
		orderRepo: orderRepo,
		timer:     timer,
		logger: config.Logger.ContextLogger.WithFields(log.Fields{
			"type": "bot",
		}),
	}, nil
}
