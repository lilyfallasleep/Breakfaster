package autoreply

// AutoReplier is the AI chatbot type
type AutoReplier struct {
	secretKey  string
	builderURL string
}

// Payload is the data payload type
type Payload struct {
	Description string `json:"description"`
}

// Bubble is the payload wrapper type
type Bubble struct {
	Type string  `json:"type"`
	Data Payload `json:"data"`
}

// ClovaRequest is the clova chatbot api request type
type ClovaRequest struct {
	Version   string   `json:"version"`
	UserID    string   `json:"userId"`
	Timestamp int64    `json:"timestamp"`
	Bubbles   []Bubble `json:"bubbles"`
	Event     string   `json:"event"`
}

// ClovaResponse is the clova chatbot api response type
type ClovaResponse struct {
	Version   string   `json:"version"`
	UserID    string   `json:"userId"`
	SessionID string   `json:"sessionId"`
	Timestamp int64    `json:"timestamp"`
	Bubbles   []Bubble `json:"bubbles"`
	Scenario  struct {
		Name   string        `json:"name"`
		Intent []interface{} `json:"intent"`
	} `json:"scenario"`
	Entities []interface{} `json:"entities"`
	Keywords []interface{} `json:"keywords"`
	Event    string        `json:"event"`
}

type errorResponseDetail struct {
	Message  string `json:"message"`
	Property string `json:"property"`
}

// ErrorResponse type
type ErrorResponse struct {
	Message string                `json:"message"`
	Details []errorResponseDetail `json:"details"`
	// OAuth Errors
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
