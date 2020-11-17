package autoreply

import (
	c "breakfaster/config"
	"breakfaster/service/constant"
	"time"
)

func (ar *AutoReplierImpl) getPrediction(body *ClovaRequest) (string, error) {
	result, err := ar.sendRequest(body)
	if err != nil {
		return "", err
	}
	return result.Bubbles[0].Data.Description, nil
}

// Predict method predicts the intent of an user
func (ar *AutoReplierImpl) Predict(userMsg string) (string, error) {
	clovaRequest := &ClovaRequest{
		Version:   "v2",
		UserID:    constant.UserID,
		Timestamp: time.Now().Unix(),
		Bubbles: []Bubble{
			Bubble{
				Type: "text",
				Data: Payload{
					Description: userMsg,
				},
			},
		},
		Event: "send",
	}
	return ar.getPrediction(clovaRequest)
}

// NewAutoReplier is the factory for AutoReplierImpl
func NewAutoReplier(config *c.Config) AutoReplier {
	return &AutoReplierImpl{
		secretKey:  config.ClovaSecretKey,
		builderURL: config.ClovaBuilderURL,
	}
}
