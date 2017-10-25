package server

import (
	"net/http"
	"log"
)

func Start(addr string) {
	http.HandleFunc("/crawl", handleCrawl)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
