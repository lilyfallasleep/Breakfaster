package autoreply

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func closeResponse(res *http.Response) error {
	defer res.Body.Close() // type of res.Body: io.ReadCloser
	// Discard is an io.Writer on which all Write calls succeed without doing anything
	_, err := io.Copy(ioutil.Discard, res.Body)
	return err
}

func isSuccess(code int) bool {
	return code/100 == 2
}

func checkResponse(res *http.Response) error {
	if isSuccess(res.StatusCode) {
		return nil
	}
	decoder := json.NewDecoder(res.Body)
	result := ErrorResponse{}
	if err := decoder.Decode(&result); err != nil {
		return &APIError{
			Code: res.StatusCode,
		}
	}
	return &APIError{
		Code:     res.StatusCode,
		Response: &result,
	}
}

func decodeToClovaResponse(res *http.Response) (*ClovaResponse, error) {
	if err := checkResponse(res); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	var result ClovaResponse
	if err := decoder.Decode(&result); err != nil {
		if err == io.EOF {
			return &result, nil
		}
		return nil, err
	}
	return &result, nil
}
