package bus

import (
	"database/sql"
	"errors"
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

	if bus.Model.CountPlace != 0 {
		err = insertModels(tx, bus.Model)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	} else {
		if exists, err := existsModel(tx, bus.Model); !exists {
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
			return fmt.Errorf("%s: %w", op, "Not found model")
		}
	}

	query := `INSERT INTO buses (state_number, model) VALUES ($1, $2)`
	_, err = tx.Exec(query, bus.StateNumber, bus.Model.Model)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	return tx.Commit()
}

func insertModels(tx *sqlx.Tx, model models.Model) error {
	query := `INSERT INTO models (model, count_places) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := tx.Exec(query, model.Model, model.CountPlace)
	if err != nil {
		return fmt.Errorf("%s: %w", query, err)
	}
	return nil
}

func existsModel(tx *sqlx.Tx, model models.Model) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM models WHERE model = $1)`
	var exists bool
	err := tx.QueryRow(query, model.Model).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("%s: %w", query, err)
	}
	return exists, nil
}

func (s *Bus) GetAll() ([]models.Bus, error) {
	const op = "storage.bus.GetAll"

	query := `
	SELECT
    	buses.state_number, m.model AS "models.model", m.count_places AS "models.count_places"
	FROM buses
	JOIN models m ON buses.model = m.model`

	var buses []models.Bus
	err := s.db.Select(&buses, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return buses, nil
}

func (s *Bus) Get(stateNumber string) (models.Bus, error) {
	const op = "storage.bus.Get"

	query := `
	SELECT 
		state_number, m.model AS "models.model", m.count_places AS "models.count_places"
	FROM buses
	JOIN models m ON buses.model = m.model
	WHERE state_number = $1`

	var bus models.Bus
	err := s.db.Get(&bus, query, stateNumber)
	if err != nil {
		return bus, fmt.Errorf("%s: %w", op, err)
	}
	return bus, nil
}
