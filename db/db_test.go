package db

import (
	"testing"
	"fmt"
)

func TestDbConnect(t *testing.T) {
	Connect(TEST_DB_CONNECT)
	defer SafeClose()

	_, err := MySQL.Exec(`INSERT INTO person (name, phone) VALUES ("golang", 123456), ("golang1", 123456)`)
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := MySQL.Query("SELECT * FROM person")
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var phone int
		var name string
		if err := rows.Scan(&name, &phone); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%s : %d\n", name, phone)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
	}

	_, err = MySQL.Exec("DELETE  FROM person")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestDbWConnect(t *testing.T) {
	Connect(DB_CONNECT)
	defer SafeClose()

	rows, err := MySQL.Query("SELECT cid,tag FROM wallpager")
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var Tag string
		var Cid string
		if err := rows.Scan(&Cid, &Tag); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("%s : %s\n", Cid, Tag)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
}
