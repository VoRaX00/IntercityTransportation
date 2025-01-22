package flight

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"kursachDB/internal/domain/models"
)

type Flight struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Flight {
	return &Flight{
		db: db,
	}
}

func (r *Flight) Add(flight models.Flight) error {
	const op = `storage.flight.Add`
	query := `INSERT INTO 
    flights (from_id, to_id, departure, arrival, state_number)
    VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query, flight.From.Id, flight.To.Id, flight.Departure, flight.Arrival, flight.Bus.StateNumber)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *Flight) GetAll() ([]models.Flight, error) {
	const op = `storage.flight.GetAll`
	query := `SELECT 
    id, from_id, to_id, departure, arrival, state_number
    FROM flights`
	var flights []models.Flight
	err := r.db.Select(&flights, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return flights, nil
}
