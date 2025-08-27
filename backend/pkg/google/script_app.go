package google

import (
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
