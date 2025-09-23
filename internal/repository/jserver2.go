package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"flight-aggregator/internal/model"
	"net/http"
	"strconv"
	"time"
)

type JServer2Repo struct { url string }
func NewJServer2Repo(url string) *JServer2Repo { return &JServer2Repo{url: url} }
func (r *JServer2Repo) Name() string { return "j-server2" }

type js2Flight struct {
	Code   string `json:"code"`
	Cost   string `json:"cost"`
	Cur    string `json:"cur"`
	DepISO string `json:"dep_time"`
	ArrISO string `json:"arr_time"`
}

func (r *JServer2Repo) Fetch(ctx context.Context) ([]model.Flight, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, r.url, nil)
	if err != nil { return nil, err }
	res, err := http.DefaultClient.Do(req)
	if err != nil { return nil, fmt.Errorf("%s: %w", r.Name(), err) }
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK { return nil, fmt.Errorf("%s: unexpected status %d", r.Name(), res.StatusCode) }

	var raw []js2Flight
	if err := json.NewDecoder(res.Body).Decode(&raw); err != nil { return nil, err }

	out := make([]model.Flight, 0, len(raw))
	for _, it := range raw {
		price, err := strconv.ParseFloat(it.Cost, 64); if err != nil { return nil, err }
		dep, err := time.Parse(time.RFC3339, it.DepISO); if err != nil { return nil, err }
		arr, err := time.Parse(time.RFC3339, it.ArrISO); if err != nil { return nil, err }
		out = append(out, model.Flight{ Provider: r.Name(), ID: it.Code, Price: price, Currency: it.Cur, Departure: dep, Arrival: arr, TravelDuration: arr.Sub(dep), })
	}
	return out, nil
}
