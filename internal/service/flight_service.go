package service

import (
	"context"
	"errors"
	"flight-aggregator/internal/model"
	"flight-aggregator/internal/repository"
	"flight-aggregator/internal/sort"
	"log"
	"sync"
	"time"
)

type FlightService struct {
	repos   []repository.FlightRepository
	timeout time.Duration
}

func NewFlightService(repos []repository.FlightRepository, timeout time.Duration) *FlightService {
	return &FlightService{repos: repos, timeout: timeout}
}

func (s *FlightService) GetFlights(ctx context.Context, sortBy sort.By, order sort.Order) ([]model.Flight, error) {
	if len(s.repos) == 0 {
		return nil, errors.New("no repositories configured")
	}
	cctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	type result struct{ flights []model.Flight; err error }
	ch := make(chan result, len(s.repos))

	var wg sync.WaitGroup
	for _, repo := range s.repos {
		wg.Add(1)
		go func(r repository.FlightRepository) {
			defer wg.Done()
			flights, err := r.Fetch(cctx)
			if err != nil { log.Printf("WARN: %s fetch failed: %v", r.Name(), err) }
			ch <- result{flights: flights, err: err}
		}(repo)
	}
	wg.Wait()
	close(ch)

	all := make([]model.Flight, 0)
	for res := range ch {
		if res.err == nil { all = append(all, res.flights...) }
	}
	sort.SortFlights(all, sortBy, order)
	return all, nil
}
