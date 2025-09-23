package controller

import (
	"encoding/json"
	"flight-aggregator/internal/model"
	"flight-aggregator/internal/service"
	"flight-aggregator/internal/sort"
	"net/http"
)

func FlightsHandler(svc *service.FlightService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		by := sort.ParseBy(r.URL.Query().Get("sort_by"))
		order := sort.ParseOrder(r.URL.Query().Get("order"))
		flights, err := svc.GetFlights(r.Context(), by, order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := make([]model.FlightResponse, 0, len(flights))
		for _, f := range flights {
			resp = append(resp, f.ToResponse())
		}
		_ = json.NewEncoder(w).Encode(resp)
	}
}
