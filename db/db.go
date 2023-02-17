package db

import (
	"database/sql"
	"log"
)

type DB struct {
	Database *sql.DB
}

func (d *DB) InitDB() {
	db, err := sql.Open("sqlite3", "/data/database/database.db")
	if err != nil {
		log.Fatal(err)
	}

	d.Database = db

	initStatement := `
		CREATE TABLE IF NOT EXISTS MANHWAS (ID  NOT NULL PRIMARY KEY, NAME VARCHAR(255) NOT NULL, PICTURE VARCHAR(255) NOT NULL)
	`

	_, err = d.Database.Exec(initStatement)
	if err != nil {
		log.Fatalf("%q: %s\n", err, initStatement)
	}

}
