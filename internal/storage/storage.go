package storage

import (
	"github.com/jmoiron/sqlx"
	authrepo "kursachDB/internal/storage/auth"
	ticketrepo "kursachDB/internal/storage/ticket"
)

type Storage struct {
	Auth   *authrepo.Auth
	Ticket *ticketrepo.Ticket
}

func New(db *sqlx.DB) *Storage {
	return &Storage{
		Auth:   authrepo.New(db),
		Ticket: ticketrepo.New(db),
	}
}
