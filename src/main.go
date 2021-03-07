package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/otonnesen/paniproject/db"
)

var DB *sql.DB

func main() {

	DB = db.GetDatabase("./links.db")

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
		log.Printf("$PORT not set, defaulting to %v", port)
	}

	http.HandleFunc("/v1/shorten", shortenHandler)

	http.ListenAndServe(":"+port, nil)
}
