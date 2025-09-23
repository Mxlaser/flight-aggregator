package model

import "time"

type Flight struct {
	Provider       string
	ID             string
	Price          float64
	Currency       string
	Departure      time.Time
	Arrival        time.Time
	TravelDuration time.Duration
}

type FlightResponse struct {
	Provider          string  `json:"provider"`
	ID                string  `json:"id"`
	Price             float64 `json:"price"`
	Currency          string  `json:"currency"`
	DepartureDate     string  `json:"departure_date"`
	ArrivalDate       string  `json:"arrival_date"`
	TimeTravelMinutes int     `json:"time_travel_minutes"`
}

func (f Flight) ToResponse() FlightResponse {
	return FlightResponse{
		Provider:          f.Provider,
		ID:                f.ID,
		Price:             f.Price,
		Currency:          f.Currency,
		DepartureDate:     f.Departure.Format(time.RFC3339),
		ArrivalDate:       f.Arrival.Format(time.RFC3339),
		TimeTravelMinutes: int(f.TravelDuration.Minutes()),
	}
}
