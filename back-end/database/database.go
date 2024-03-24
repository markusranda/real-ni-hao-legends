package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() error {
	log.Println("Initiliazing database")

	// Open a connection to the SQLite database.
	db, err := sql.Open("sqlite", "nihao.db")
	if err != nil {
		return err
	}

	// Create the states table if it doesn't already exist.
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS states (
			id INTEGER PRIMARY KEY,
			state_data TEXT
		);`)

	if err != nil {
		panic(fmt.Sprintf("Couldn't create the most important table: %s", err.Error()))
	}

	DB = db
	return nil
}
