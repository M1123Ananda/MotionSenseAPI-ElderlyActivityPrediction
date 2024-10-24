package utils

import (
	"bytes"
	"net/http"
)

func CallPostRequest(url string, jsonData []byte) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
