package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"flight-aggregator/internal/model"
	"net/http"
	"time"
)

type JServer1Repo struct { url string }
func NewJServer1Repo(url string) *JServer1Repo { return &JServer1Repo{url: url} }
func (r *JServer1Repo) Name() string { return "j-server1" }

type js1Flight struct {
	ID        string  `json:"id"`
	Price     float64 `json:"price"`
	Currency  string  `json:"currency"`
	Departure string  `json:"departure"`
	Arrival   string  `json:"arrival"`
}

func (r *JServer1Repo) Fetch(ctx context.Context) ([]model.Flight, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.url, nil)
	if err != nil { return nil, err }
	res, err := http.DefaultClient.Do(req)
	if err != nil { return nil, fmt.Errorf("%s: %w", r.Name(), err) }
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK { return nil, fmt.Errorf("%s: unexpected status %d", r.Name(), res.StatusCode) }

	var raw []js1Flight
	if err := json.NewDecoder(res.Body).Decode(&raw); err != nil { return nil, err }

	out := make([]model.Flight, 0, len(raw))
	for _, it := range raw {
		dep, err := time.Parse(time.RFC3339, it.Departure); if err != nil { return nil, err }
		arr, err := time.Parse(time.RFC3339, it.Arrival); if err != nil { return nil, err }
		out = append(out, model.Flight{ Provider: r.Name(), ID: it.ID, Price: it.Price, Currency: it.Currency, Departure: dep, Arrival: arr, TravelDuration: arr.Sub(dep), })
	}
	return out, nil
}
