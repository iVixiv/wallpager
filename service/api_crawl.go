package server

import (
	"net/http"
	"encoding/json"
	"fmt"
	"wallpager/crawler"
)

func handleCrawl(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("content-type", "text/html; charset=utf-8")
	defer r.Body.Close()

	Response(w, HandlerSuccessResult(nil, crawler.Crawl_Type(10000)))
}

func HandlerSuccessResult(data interface{}, err error) (string) {
	result := Result{}
	if err == nil {
		result.Code = 200
		result.Message = "success"
		result.Data = data
	} else {
		result.Code = 500
		result.Message = err.Error()
	}
	responseBody, _ := json.Marshal(result)
	return string(responseBody)
}

func Response(w http.ResponseWriter, data string) {
	fmt.Fprint(w, data)
}

type Result struct {
	Code    int
	Message string
	Data    interface{}
}
