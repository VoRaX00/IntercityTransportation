package mapper

import (
	"fmt"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
	"time"
)

func PlaceAddToPlace(place services.AddPlace) models.Place {
	return models.Place{
		NamePlace: place.Name,
		Type: models.TypePlace{
			Type: place.Type,
		},
	}
}

func FlightAddToFlight(flight services.AddFlight) (models.Flight, error) {
	departure, err := time.Parse("02.01.2006", flight.Departure)
	if err != nil {
		return models.Flight{}, fmt.Errorf("wrong departure")
	}

	arrival, err := time.Parse("02.01.2006", flight.Arrival)
	if err != nil {
		return models.Flight{}, fmt.Errorf("wrong arrival")
	}
	return models.Flight{
		From: models.Place{
			NamePlace: flight.From,
		},
		To: models.Place{
			NamePlace: flight.To,
		},
		Departure: departure,
		Arrival:   arrival,
		Bus: models.Bus{
			StateNumber: flight.StateNumber,
		},
	}, nil
}
