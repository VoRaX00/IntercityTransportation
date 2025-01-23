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
	tx, err := r.db.Beginx()

	fromId, err := getPlace(tx, flight.From.NamePlace)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	toId, err := getPlace(tx, flight.To.NamePlace)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	query := `INSERT INTO 
    flights (from_id, to_id, departure, arrival, state_number)
    VALUES ($1, $2, $3, $4, $5)`

	_, err = tx.Exec(query, fromId, toId, flight.Departure, flight.Arrival, flight.Bus.StateNumber)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	return tx.Commit()
}

func getPlace(tx *sqlx.Tx, place string) (int, error) {
	query := `SELECT id FROM places WHERE name_place = $1`
	var placeId int
	if err := tx.QueryRow(query, place).Scan(&placeId); err != nil {
		return 0, fmt.Errorf("%s: %w", query, err)
	}
	return placeId, nil
}

func (r *Flight) GetAll() ([]models.Flight, error) {
	const op = `storage.flight.GetAll`
	query := `SELECT 
    f.id,
    p_from.id AS "places.id",
    p_to.id AS "to_places.id",  
    departure, arrival,
    b.state_number AS "buses.state_number"
    FROM flights f
    JOIN places p_from ON p_from.id = f.from_id
    JOIN places p_to ON p_to.id = f.to_id
    JOIN buses b ON b.state_number = f.state_number`

	var flights []models.Flight
	err := r.db.Select(&flights, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return flights, nil
}
