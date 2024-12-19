package http

import (
	"io"
	"net/http"
)

func Get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	body := string(data)
	return body, nil
}
