package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
		log.Printf("$PORT not set, defaulting to %v", port)
	}

	// POST: Shortens a URL.
	//	Request:
	//		{
	//		  "long_url": "string"
	//		}
	//	Response:
	//		{
	//		  "hash": "string"
	//		}
	http.HandleFunc("/v1/shorten", shortenHandler)
	// GET: Retrieves original URL corresponding to the hash in
	// `/v1/hash/{hash}`.
	//	Response:
	//		{
	//		  "long_url": "string"
	//		}
	http.HandleFunc("/v1/hash/", hashHandler)

	http.ListenAndServe(":"+port, nil)
}
