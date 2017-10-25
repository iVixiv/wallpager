package server

import (
	"net/http"
	"wallpager/wallpager"
)

func handleWallPager(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("content-type", "text/html; charset=utf-8")
	defer r.Body.Close()
	r.ParseForm()

	Response(w, HandlerSuccessResult(wallpager.Select(r.Form.Get("start"), r.Form.Get("type"))))
}
