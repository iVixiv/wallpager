package main

import (
	"github.com/op/go-logging"
	"fmt"
	"wallpager/service"
	"wallpager/db"
)

var log = logging.MustGetLogger("main")

func main() {
	db.Connect(db.DB_CONNECT)
	defer db.SafeClose()

	addr := fmt.Sprintf(":%s", db.SERVICE_PORT)
	log.Info("Server Listening %s", addr)

	server.Start(addr)
}
