package communication

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func RequestJSON(method string, url string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Internal-Call", "true")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &http.Response{}, err
	}
	defer resp.Body.Close()

	return resp, nil
}

// PostJSON sends a POST request with JSON data
func PostJSON(url string, data []byte) (*http.Response, error) {
	return RequestJSON("POST", url, data)
}

// ParseJSONResponse parses a JSON response
func ParseJSONResponse(data http.Response) (interface{}, error) {
	var result interface{}
	err := json.NewDecoder(data.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
