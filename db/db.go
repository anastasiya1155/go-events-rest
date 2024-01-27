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
	createUsersTabs := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUsersTabs)

	if err != nil {
		panic("Migration failed!")
	}

	var createEventsTable = `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	stmt, err := DB.Prepare(createEventsTable)

	if err != nil {
		fmt.Println()
		panic("Create events table query invalid")
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	if err != nil {
		panic("Migration failed!")
	}
}
