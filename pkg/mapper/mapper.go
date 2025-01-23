package mapper

import (
	"fmt"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
	"strings"
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
	flight.Departure = strings.TrimSpace(flight.Departure)
	flight.Arrival = strings.TrimSpace(flight.Arrival)
	departure, err := time.Parse("02.01.2006 15:04", flight.Departure)
	if err != nil {
		return models.Flight{}, fmt.Errorf("wrong departure")
	}

	arrival, err := time.Parse("02.01.2006 15:04", flight.Arrival)
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

func BuyTicketToTicket(ticket services.BuyTicket) models.Ticket {
	var t models.Ticket
	t.Cost = ticket.Cost
	t.Flight = make([]models.Flight, len(ticket.Flights))
	for i, val := range ticket.Flights {
		t.Flight[i] = models.Flight{
			Id: val,
		}
	}

	t.User = models.User{
		PhoneNumber: ticket.PhoneNumber,
	}
	return t
}
