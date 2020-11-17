package messaging

import "breakfaster/mybot"

// MessageControllerImpl implements LINE bot messgae pushing functionality
type MessageControllerImpl struct {
	pusher *mybot.BreakFastPusher
}

// NewMessageController is a factory for MessageControllerImpl
func NewMessageController(pusher *mybot.BreakFastPusher) MessageController {
	return &MessageControllerImpl{
		pusher: pusher,
	}
}
