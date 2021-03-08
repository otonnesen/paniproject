// The db_worker module extracts any database transactions from hash_worker
// into its own program so the hash_workers can be scaled out and still use a
// sqlite database. This program would be replaced by a regular database server.
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/otonnesen/paniproject/db_server/db"
)

var DB *sql.DB

func main() {

	DB = db.GetDatabase("./links.db")

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
		log.Printf("$PORT not set, defaulting to %v", port)
	}

	// POST: Creates a hash for the specified url and adds a new entry
	// into the links table.
	//	Request:
	//		{
	//		  "url": "string"
	//		}
	//	Response:
	//		{
	//		  "id": number
	//		}
	http.HandleFunc("/v1/link", linkHandler)
	// GET: Retrieves the original URL corresponding to the id in
	// '/v1/link/{id}'.
	//	Response:
	//		{
	//		  "url": "string"
	//		}
	http.HandleFunc("/v1/link/", linkIdHandler)

	http.ListenAndServe(":"+port, nil)
}
