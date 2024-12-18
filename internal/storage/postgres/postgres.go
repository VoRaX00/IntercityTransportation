package postgres

import (
	"github.com/jmoiron/sqlx"
	authrepo "kursachDB/internal/storage/auth"
	ticketrepo "kursachDB/internal/storage/ticket"
)

type Storage struct {
	db     *sqlx.DB
	Auth   *authrepo.Auth
	Ticket *ticketrepo.Ticket
}

func New(storagePath string) (*Storage, error) {
	db, err := sqlx.Open("postgres", storagePath)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db:     db,
		Auth:   authrepo.New(db),
		Ticket: ticketrepo.New(db),
	}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
