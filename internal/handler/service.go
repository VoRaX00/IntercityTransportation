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
		Bus:    bus.New(log, repos.Bus),
		Flight: flight.New(log, repos.Flight),
		Place:  place.New(log, repos.Place),
		Ticket: ticket.New(log, repos.Ticket),
	}
}
