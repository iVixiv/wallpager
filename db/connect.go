package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/op/go-logging"
)

var (
	log   = logging.MustGetLogger("db")
	MySQL *sql.DB
)

// 连接到 MySQL 数据库使用 config.MYSQL_DATA_SOURCE_NAME
// 可以设置环境变量 MYSQL_DATA_SOURCE_NAME 来更改 dataSourceName
// http://godoc.org/github.com/go-sql-driver/mysql
// 务必调用 defer db.SafeClose() 安全关闭连接
func Connect(connect string) {
	var err error
	MySQL, err = sql.Open("mysql", connect)
	if err != nil {
		log.Errorf("Failed to open mysql connect : %s", err.Error())
	}
	// 设置最大连接数
	MySQL.SetMaxOpenConns(500)
	// 设置最大闲置连接数
	MySQL.SetMaxIdleConns(100)
	err = MySQL.Ping()
	log.Info("Connected to the mysql!")
	if err != nil {
		log.Panic(err)
	}
}

// 安全关闭 MySQL 连接
func SafeClose() {
	if MySQL != nil {
		MySQL.Close()
	}
}
