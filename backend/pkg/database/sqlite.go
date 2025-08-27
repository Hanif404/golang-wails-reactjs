package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var db *sql.DB

func init() {
	dba, err := sql.Open("sqlite3", "./data.sqlite")
	if err != nil {
		panic(err)
	}
	db = dba
	createTable()
}

func ConnDB() *sql.DB {
	return db
}

func createTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS session (id INTEGER PRIMARY KEY, name TEXT, email TEXT, password TEXT, loginAt INTEGER)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS sync_sessions (id INTEGER PRIMARY KEY, syncType TEXT, createdAt INTEGER, startDate INTEGER, endDate INTEGER, message TEXT, status TEXT)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS sync_logs (id INTEGER PRIMARY KEY, sessionId INTEGER, tableName TEXT, recordId INTEGER, operation TEXT, syncDirection TEXT, status TEXT)")
	if err != nil {
		panic(err)
	}
}
