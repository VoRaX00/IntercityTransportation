package place

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"kursachDB/internal/domain/models"
)

type Place struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Place {
	return &Place{db: db}
}

func (s *Place) Add(place models.Place) error {
	const op = "storage.place.Add"
	tx, err := s.db.Beginx()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	var typeId int64
	query := `SELECT id FROM types_places WHERE type_place = $1`
	row, err := tx.Query(query, place.Type.Type)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = row.Scan(&typeId); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	query = `INSERT INTO places (name_place, type_id) VALUES ($1, $2)`
	_, err = tx.Exec(query, place.NamePlace, typeId)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Place) Delete(id int) error {
	const op = "storage.place.Delete"
	tx, err := s.db.Beginx()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	query := `DELETE FROM places WHERE id = $1`

	_, err = tx.Exec(query, id)
	if err != nil {
		_ = tx.Rollback()
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Place) GetAll() ([]models.Place, error) {
	const op = "storage.place.GetAll"
	var places []models.Place
	query := `SELECT 
    	id, name_place, types_places.id, types_places.type_place
		FROM places
		JOIN types_places ON id = places.type_id`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// TODO: implemented
	return nil, nil
}
