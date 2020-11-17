package autoreply

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (ar *AutoReplierImpl) sendRequest(body *ClovaRequest) (*ClovaResponse, error) {
	var jsonstr []byte
	var signature string
	var err error

	jsonstr, err = json.Marshal(body)
	if err != nil {
		return &ClovaResponse{}, err
	}
	signature, err = ar.generateSignature(jsonstr)
	if err != nil {
		return &ClovaResponse{}, err
	}

	req, err := http.NewRequest("POST", ar.builderURL, bytes.NewBuffer(jsonstr))
	req.Header.Set("X-NCP-CHATBOT_SIGNATURE", signature)
	req.Header.Set("Content-Type", "application/json;UTF-8")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &ClovaResponse{}, err
	}

	defer closeResponse(res) // remember to close res
	return decodeToClovaResponse(res)

}
