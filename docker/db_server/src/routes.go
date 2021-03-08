package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/otonnesen/paniproject/db_server/db"
)

func linkHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		d := struct {
			LongURL string `json:"long_url"`
		}{}
		err := json.NewDecoder(req.Body).Decode(&d)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id := db.GetNextId(DB, d.LongURL)

		log.Printf("Database entry created for id %v\n", id)
		s := struct {
			ID int `json:"id"`
		}{
			ID: id,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(s)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func linkIdHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(req.URL.Path, "/v1/link/"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	link, err := db.GetLinkById(DB, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	s := struct {
		LongURL string `json:"long_url"`
	}{
		LongURL: link.LongURL,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(s)
}
