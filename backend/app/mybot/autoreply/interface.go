package autoreply

// AutoReplier is the interface for auto reply service
type AutoReplier interface {
	Predict(userMsg string) (string, error)
}
