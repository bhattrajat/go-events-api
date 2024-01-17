package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() error {
	var err error
	DB, err = sql.Open("sqlite", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	createTable := `
	CREATE TABLE IF NOT EXISTS events 
		(
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			userId INTEGER NOT NULL
		);
	`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
