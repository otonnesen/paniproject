package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/otonnesen/paniproject/api"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
		log.Printf("$PORT not set, defaulting to %v", port)
	}

	http.HandleFunc("/v1/shorten", shortenHandler)

	http.ListenAndServe(":"+port, nil)
}

func shortenHandler(w http.ResponseWriter, req *http.Request) {
	data, err := api.NewShortenRequest(req)
	data = data
	if err != nil {
		log.Printf("Bad shorten request: %v\n", err)
	}

	// shortenResp := logic(data)
	shortenResp := ""

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shortenResp)
}

func getShortURL(url string) string {
	// Ask db what number I am
	// Generate hash (do this in stored procedure?)
	// append hash to domain (localhost)
	// Insert hash into database (unless I already did this in step 2)
	return ""
}
