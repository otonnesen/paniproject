package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func shortenHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		d := struct {
			LongURL string `json:"long_url"`
		}{}
		err := json.NewDecoder(req.Body).Decode(&d)
		if err != nil || d.LongURL == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		json_data, err := json.Marshal(d)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp, err := http.Post("http://paniproject_db_1:3000/v1/link", "application/json", bytes.NewBuffer(json_data))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r := struct {
			ID int `json:"id"`
		}{}
		err = json.NewDecoder(resp.Body).Decode(&r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("ID: %v\n", r.ID)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func hashHandler(w http.ResponseWriter, req *http.Request) {
}
