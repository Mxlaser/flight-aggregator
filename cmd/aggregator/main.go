package main

import (
	"log"
	"flight-aggregator/internal/config"
	"flight-aggregator/internal/httpserver"
	"flight-aggregator/internal/repository"
	"flight-aggregator/internal/service"
)

func main() {
	cfg := config.Load()
	repo1 := repository.NewJServer1Repo(cfg.JServer1URL)
	repo2 := repository.NewJServer2Repo(cfg.JServer2URL)
	svc := service.NewFlightService([]repository.FlightRepository{repo1, repo2}, cfg.RequestTimeout)
	srv := httpserver.New(cfg.Addr, svc)
	log.Printf("Flight Aggregator listening on %s", cfg.Addr)
	if err := srv.ListenAndServe(); err != nil { log.Fatal(err) }
}
