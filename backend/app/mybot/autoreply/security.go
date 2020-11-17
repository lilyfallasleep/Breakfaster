package autoreply

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func (ar *AutoReplierImpl) generateSignature(requestBody []byte) (string, error) {
	hash := hmac.New(sha256.New, []byte(ar.secretKey))

	// hash body with secret key
	_, err := hash.Write(requestBody)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash.Sum(nil)), nil
}
