package tests

import (
	"flight-aggregator/internal/model"
	"flight-aggregator/internal/sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func tTS(ti string) time.Time { d, _ := time.Parse(time.RFC3339, ti); return d }

func TestSortByPriceAsc(t *testing.T) {
	fs := []model.Flight{{Price: 300},{Price: 200},{Price: 250}}
	sort.SortFlights(fs, sort.ByPrice, sort.OrderAsc)
	assert.Equal(t, 200.0, fs[0].Price)
	assert.Equal(t, 250.0, fs[1].Price)
	assert.Equal(t, 300.0, fs[2].Price)
}

func TestSortByPriceDesc(t *testing.T) {
	fs := []model.Flight{{Price: 100},{Price: 50},{Price: 200}}
	sort.SortFlights(fs, sort.ByPrice, sort.OrderDesc)
	assert.Equal(t, 200.0, fs[0].Price)
	assert.Equal(t, 100.0, fs[1].Price)
	assert.Equal(t, 50.0, fs[2].Price)
}

func TestSortByTimeTravel(t *testing.T) {
	fs := []model.Flight{{TravelDuration: 180 * time.Minute},{TravelDuration: 60 * time.Minute},{TravelDuration: 90 * time.Minute}}
	sort.SortFlights(fs, sort.ByTimeTravel, sort.OrderAsc)
	assert.Equal(t, 60*time.Minute, fs[0].TravelDuration)
	assert.Equal(t, 90*time.Minute, fs[1].TravelDuration)
	assert.Equal(t, 180*time.Minute, fs[2].TravelDuration)
}

func TestSortByDepartureDateDesc(t *testing.T) {
	fs := []model.Flight{{Departure: tTS("2025-10-01T08:00:00Z")},{Departure: tTS("2025-09-30T08:00:00Z")},{Departure: tTS("2025-10-05T08:00:00Z")}}
	sort.SortFlights(fs, sort.ByDepartureDate, sort.OrderDesc)
	assert.Equal(t, tTS("2025-10-05T08:00:00Z"), fs[0].Departure)
	assert.Equal(t, tTS("2025-10-01T08:00:00Z"), fs[1].Departure)
	assert.Equal(t, tTS("2025-09-30T08:00:00Z"), fs[2].Departure)
}
