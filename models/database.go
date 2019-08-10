package models

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var instance *sql.DB

func getDatabase() *sql.DB {

	if instance == nil {

		db, err := sql.Open("mysql", "golang:golf@/golf?parseTime=true")
		if err != nil {
			panic(err)
		}

		instance = db

	}

	return instance

}