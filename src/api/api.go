// Package api provides type definitions and decoders for API objects.
package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

type ShortenResponse struct {
	ID        int       `json:"id"`
	LongURL   string    `json:"long_url"`
	ShortURL  string    `json:"short_url"`
	CreatedAt time.Time `json:"created_at"`
}

func NewShortenRequest(req *http.Request) (*ShortenRequest, error) {
	d := ShortenRequest{}
	err := json.NewDecoder(req.Body).Decode(&d)
	return &d, err
}
