package sort

import (
	"flight-aggregator/internal/model"
	gosort "sort"
	"strings"
)

type By string
type Order string

const (
	ByPrice         By = "price"
	ByTimeTravel    By = "time_travel"
	ByDepartureDate By = "departure_date"

	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

func ParseBy(s string) By {
	switch strings.ToLower(s) {
	case string(ByPrice), "":
		return ByPrice
	case string(ByTimeTravel):
		return ByTimeTravel
	case string(ByDepartureDate):
		return ByDepartureDate
	default:
		return ByPrice
	}
}

func ParseOrder(s string) Order {
	switch strings.ToLower(s) {
	case string(OrderDesc):
		return OrderDesc
	default:
		return OrderAsc
	}
}

func SortFlights(flights []model.Flight, by By, order Order) {
	less := func(i, j int) bool {
		switch by {
		case ByTimeTravel:
			if order == OrderAsc { return flights[i].TravelDuration < flights[j].TravelDuration }
			return flights[i].TravelDuration > flights[j].TravelDuration
		case ByDepartureDate:
			if order == OrderAsc { return flights[i].Departure.Before(flights[j].Departure) }
			return flights[i].Departure.After(flights[j].Departure)
		case ByPrice:
			fallthrough
		default:
			if order == OrderAsc { return flights[i].Price < flights[j].Price }
			return flights[i].Price > flights[j].Price
		}
	}
	gosort.SliceStable(flights, less)
}
