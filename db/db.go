package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("sqlite3", "todos.db")
	if err != nil {
		log.Fatalf("error opening database.\n%v", err.Error())
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		importance TEXT NOT NULL CHECK (importance IN ('low', 'mid', 'high')),
		completed BOOLEAN NOT NULL DEFAULT 0,
		date DATE NOT NULL
	);`)
	if err != nil {
		log.Fatalf("error creating todos table.\n%v", err.Error())
	}
}

func CloseDB() {
	err = DB.Close()
	if err != nil {
		log.Fatalf("error closing database.\n%v", err.Error())
	}
}
