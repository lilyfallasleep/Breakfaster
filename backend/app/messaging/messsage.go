package messaging

import "breakfaster/mybot"

// Message provides LINE bot messgae pushing functionality
type Message struct {
	pusher *mybot.BreakFastPusher
}

// NewMessage is a factory for message instance
func NewMessage(pusher *mybot.BreakFastPusher) *Message {
	return &Message{
		pusher: pusher,
	}
}
