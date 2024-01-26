package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("DB connection failed")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	var createEventsTable = `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)`
	_, err := DB.Prepare(createEventsTable)

	if err != nil {
		fmt.Println()
		panic("Create table query invalid")
	}
	// defer stmt.Close()

	// _, err = stmt.Exec()

	// if err != nil {
	// 	panic("Migration failed!")
	// }
}
