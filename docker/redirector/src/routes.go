package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func redirectHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		hashStr := strings.TrimPrefix(req.URL.Path, "/v1/hash/")

		resp, err := http.Get(fmt.Sprintf("http://nginx:5000/v1/hash/%v", hashStr))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 page not found")
			return
		}
		respData := struct {
			URL string `json:"url"`
		}{}
		err = json.NewDecoder(resp.Body).Decode(&respData)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404 page not found")
			return
		}
		url := respData.URL
		if !(strings.HasPrefix(respData.URL, "http://") ||
			strings.HasPrefix(respData.URL, "https://")) {
			url = "http://" + url
		}
		http.Redirect(w, req, url, http.StatusMovedPermanently)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
