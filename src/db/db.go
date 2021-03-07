// Package db provides types and functions for transacting with the database.
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// LinksEntry corresponds to a row in the links table in the database.
type LinksEntry struct {
	ID        int
	LongURL   string
	Hash      string
	CreatedAt time.Time
}

// GetDatabase creates the sqlite database file if it doesn't already exist,
// then opens it and returns the connection.
func GetDatabase(path string) *sql.DB {
	var f *os.File
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err = os.Create(path)
		if err != nil {
			log.Printf("Error creating database file: %v\n", err)
		}
		f.Close()
		log.Printf("%v created\n", path)
	}

	sqliteDB, _ := sql.Open("sqlite3", path)
	InitDatabase(sqliteDB)

	return sqliteDB
}

// InitDatabase creates the links table if it doesn't exist.
func InitDatabase(db *sql.DB) {
	createLinksTableSQL := `CREATE TABLE IF NOT EXISTS links (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"long_url" TEXT,
		"hash" TEXT UNIQUE,
		"created_at" TEXT
	);`

	statement, err := db.Prepare(createLinksTableSQL)
	if err != nil {
		log.Printf("Error creating links table: %v\n", err)
	}
	statement.Exec()
}

// GetNextId inserts a new row into the links table (to be filled with a hash
// at a later time) and returns the value assigned to its id column.
func GetNextId(db *sql.DB) int {
	getNextIdSQL := `INSERT INTO links (created_at) VALUES (?);`

	statement, err := db.Prepare(getNextIdSQL)
	if err != nil {
		log.Printf("Error inserting into links table: %v\n", err)
	}
	r, err := statement.Exec(time.Now().Format(time.RFC3339))
	if err != nil {
		log.Printf("Error executing statement: %v\n", err)
	}

	id, err := r.LastInsertId()
	if err != nil {
		log.Printf("Error getting last ID: %v\n", err)
	}
	return int(id)
}

// InsertLink completes a half-filled row in the links table by assigning
// values to its long_url and hash columns.
func InsertLink(db *sql.DB, id int, longURL string, hash string) {
	insertLinkSQL := `UPDATE links SET long_url = ?, hash = ? WHERE id = ?`
	statement, err := db.Prepare(insertLinkSQL)
	if err != nil {
		log.Printf("Error updating links table: %v\n", err)
	}
	_, err = statement.Exec(longURL, hash, id)
	if err != nil {
		log.Printf("Error executing statement: %v\n", err)
	}
}

// GetLinkById returns the row in the links table corresponding to the
// specified id.
func GetLinkById(db *sql.DB, id int) LinksEntry {
	row := db.QueryRow(fmt.Sprintf("SELECT * FROM links WHERE id = %v", id))

	var e LinksEntry
	err := row.Scan(&e.ID, &e.LongURL, &e.Hash, &e.CreatedAt)
	if err != nil {
		log.Printf("Error querying links: %v\n", err)
	}
	return e
}
