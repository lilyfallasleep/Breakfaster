package messaging

// MessageController is the interface for LINE bot messgae pushing functionality
type MessageController interface {
	BroadCastMenu() error
}
