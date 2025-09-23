package tests

import (
	"context"
	"flight-aggregator/internal/model"
	"flight-aggregator/internal/repository"
	"flight-aggregator/internal/service"
	"flight-aggregator/internal/sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	name    string
	flights []model.Flight
	err     error
}
func (m mockRepo) Name() string { return m.name }
func (m mockRepo) Fetch(ctx context.Context) ([]model.Flight, error) { return m.flights, m.err }

func TestFlightService_MergeAndSort(t *testing.T) {
	r1 := mockRepo{name: "r1", flights: []model.Flight{{Provider: "r1", ID: "A", Price: 150},{Provider: "r1", ID: "B", Price: 50}}}
	r2 := mockRepo{name: "r2", flights: []model.Flight{{Provider: "r2", ID: "C", Price: 100}}}

	repos := []repository.FlightRepository{r1, r2}
	svc := service.NewFlightService(repos, 2*time.Second)

	got, err := svc.GetFlights(context.Background(), sort.ByPrice, sort.OrderAsc)
	assert.NoError(t, err)
	assert.Len(t, got, 3)
	assert.Equal(t, "B", got[0].ID)
	assert.Equal(t, "C", got[1].ID)
	assert.Equal(t, "A", got[2].ID)
}
