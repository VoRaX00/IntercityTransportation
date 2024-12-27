package postgres

import (
	"github.com/jmoiron/sqlx"
	authrepo "kursachDB/internal/storage/postgres/auth"
	busrepo "kursachDB/internal/storage/postgres/bus"
	flightrepo "kursachDB/internal/storage/postgres/flight"
	placerepo "kursachDB/internal/storage/postgres/place"
	ticketrepo "kursachDB/internal/storage/postgres/ticket"
)

type Storage struct {
	db     *sqlx.DB
	Auth   *authrepo.Auth
	Bus    *busrepo.Bus
	Flight *flightrepo.Flight
	Place  *placerepo.Place
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
		Bus:    busrepo.New(db),
		Flight: flightrepo.New(db),
		Place:  placerepo.New(db),
		Ticket: ticketrepo.New(db),
	}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
