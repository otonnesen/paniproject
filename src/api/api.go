package api

import (
	"encoding/json"
	"net/http"
)

type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

func NewShortenRequest(req *http.Request) (*ShortenRequest, error) {
	d := ShortenRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}
