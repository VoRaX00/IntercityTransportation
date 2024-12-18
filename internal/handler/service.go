package handler

import (
	"kursachDB/internal/services/auth"
	"kursachDB/internal/services/ticket"
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
		Auth:   auth.New(log, repos.Auth),
		Ticket: ticket.New(log, repos.Ticket),
	}
}
