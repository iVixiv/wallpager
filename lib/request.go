package lib

import (
	"io"
	"net/http"
)

func Request(url string, method string, headerParams map[string]string, body io.Reader, charType string) (string, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}

	for k, v := range headerParams {
		req.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := DecodeReader(res.Body, charType)
	if err != nil {
		return "", err
	}
	return string(resBody), nil
}
