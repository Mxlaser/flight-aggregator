package httpserver

import (
	"net/http"

	"flight-aggregator/internal/controller"
	"flight-aggregator/internal/service"
)

func New(addr string, svc *service.FlightService) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", controller.HealthHandler)
	mux.HandleFunc("/flight", controller.FlightsHandler(svc))
	return &http.Server{ Addr: addr, Handler: withJSONHeaders(mux) }
}

func withJSONHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
