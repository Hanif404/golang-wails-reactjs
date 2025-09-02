package google

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

var url = "https://script.google.com/macros/s/"

func scriptApp(key string) string {
	return url + key + "/exec"
}

func GetResponse(key, parameter string) ([]byte, error) {
	url := scriptApp(key)
	resp, err := http.Get(url + parameter)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func PostResponse(key, parameter string, payload any) ([]byte, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	url := scriptApp(key)
	req, err := http.NewRequest("POST", url+parameter, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}
	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Create HTTP client and send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
