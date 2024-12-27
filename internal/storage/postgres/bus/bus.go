package bus

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"kursachDB/internal/domain/models"
)

type Bus struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Bus {
	return &Bus{db: db}
}

func (s *Bus) Add(bus models.Bus) error {
	const op = "storage.bus.Add"
	tx, err := s.db.Beginx()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = insertModels(tx, bus.Model)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	query := `INSERT INTO buses (state_number, model) VALUES ($1, $2)`
	_, err = tx.Exec(query, bus.StateNumber, bus.Model.Model)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	return tx.Commit()
}

func insertModels(tx *sqlx.Tx, models models.Model) error {
	query := `INSERT INTO models (model, count_places) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := tx.Exec(query, models)
	if err != nil {
		return fmt.Errorf("%s: %w", query, err)
	}
	return nil
}

func (s *Bus) Delete(stateNumber string) error {
	const op = "storage.bus.Delete"
	query := `DELETE FROM buses WHERE state_number = $1`

	_, err := s.db.Exec(query, stateNumber)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Bus) GetAll() ([]models.Bus, error) {
	const op = "storage.bus.GetAll"

	query := `
	SELECT
    	buses.state_number, buses.model, models.count_places
	FROM buses
	JOIN models ON buses.model = models.model`

	var buses []models.Bus
	err := s.db.Select(&buses, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return buses, nil
}
