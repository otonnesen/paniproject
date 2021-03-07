package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/otonnesen/paniproject/api"
	"github.com/otonnesen/paniproject/db"
	"github.com/otonnesen/paniproject/hash"
)

func shortenHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		data, err := api.NewShortenRequest(req)
		data = data
		if err != nil {
			log.Printf("Bad shorten request: %v\n", err)
		}

		id := db.GetNextId(DB)
		hash, err := hash.HashFromNumber(id)
		if err != nil {
			log.Printf("Error hashing: %v\n", err)
		}

		// TODO: put this somewhere else
		DOMAIN := "localhost:5000"

		shortURL := DOMAIN + "/" + hash
		s := api.ShortenResponse{
			ID:        id,
			LongURL:   data.LongURL,
			ShortURL:  shortURL,
			CreatedAt: time.Now(),
		}

		db.InsertLink(DB, s.ID, s.LongURL, s.ShortURL)
		// entry, _ := DB.GetLink(s.ID)
		log.Printf("Database entry created for id %v\n", s.ID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(s)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
