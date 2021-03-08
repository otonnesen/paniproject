package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/otonnesen/paniproject/web/hash"
)

func shortenHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		reqData := struct {
			URL string `json:"url"`
		}{}
		err := json.NewDecoder(req.Body).Decode(&reqData)
		if err != nil || reqData.URL == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		jsonData, err := json.Marshal(reqData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resp, err := http.Post(
			"http://db:3000/v1/link",
			"application/json",
			bytes.NewBuffer(jsonData),
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		respData := struct {
			ID int `json:"id"`
		}{}
		err = json.NewDecoder(resp.Body).Decode(&respData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		hashStr, err := hash.Hash(respData.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Printf("Hash %v created for URL %v.\n", hashStr, reqData.URL)

		r := struct {
			Hash string `json:"hash"`
		}{
			Hash: hashStr,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func hashHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		hashStr := strings.TrimPrefix(req.URL.Path, "/v1/hash/")

		id, err := hash.Unhash(hashStr)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		resp, err := http.Get(fmt.Sprintf("http://db:3000/v1/link/%v", id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		respData := struct {
			URL string `json:"url"`
		}{}
		err = json.NewDecoder(resp.Body).Decode(&respData)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(respData)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
