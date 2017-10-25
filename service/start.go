package server

import (
	"net/http"
	"log"
)

func Start(addr string) {
	http.HandleFunc("/crawl", handleCrawl)
	http.HandleFunc("/wallpager", handleWallPager)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
