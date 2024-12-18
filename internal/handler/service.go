package handler

import (
	"kursachDB/internal/services/auth"
	"kursachDB/internal/services/place"
	"kursachDB/internal/services/schedule"
	"kursachDB/internal/services/ticket"
	"kursachDB/internal/services/transport"
	"kursachDB/internal/storage"
	"log/slog"
)

type Service struct {
	Auth      Auth
	Place     Place
	Schedule  Schedule
	Transport Transport
	Ticket    Ticket
}

func NewService(log *slog.Logger, repos *storage.Storage) *Service {
	return &Service{
		Auth:      auth.New(log, repos.Auth),
		Place:     place.New(log, repos),
		Schedule:  schedule.New(log, repos),
		Transport: transport.New(log, repos),
		Ticket:    ticket.New(log, repos.Ticket),
	}
}
