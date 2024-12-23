package handler

import (
	"kursachDB/internal/services/auth"
	"kursachDB/internal/services/bus"
	"kursachDB/internal/services/flight"
	"kursachDB/internal/services/place"
	"kursachDB/internal/services/ticket"
	"kursachDB/internal/storage/postgres"
	"log/slog"
)

type Service struct {
	Auth   Auth
	Place  Place
	Flight Flight
	Bus    Bus
	Ticket Ticket
}

func NewService(log *slog.Logger, repos *postgres.Storage) *Service {
	return &Service{
		Auth:   auth.New(log, repos.Auth),
		Place:  place.New(log, repos),
		Flight: flight.New(log, repos),
		Bus:    bus.New(log, repos),
		Ticket: ticket.New(log, repos.Ticket),
	}
}
