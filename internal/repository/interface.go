package repository

import (
	"context"
	"flight-aggregator/internal/model"
)

type FlightRepository interface {
	Name() string
	Fetch(ctx context.Context) ([]model.Flight, error)
}
