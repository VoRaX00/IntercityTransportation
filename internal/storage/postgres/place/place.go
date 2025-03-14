package place

import (
	"database/sql"
	"errors"
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

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		}
	}()

	exists, err := existsType(tx, place.Type.Type)
	if !exists {
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
		err = insertType(tx, place)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	var typeId int64
	query := `SELECT id FROM types_places WHERE type_place = $1`
	err = tx.QueryRow(query, place.Type.Type).Scan(&typeId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%s: %w", op, err)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	query = `INSERT INTO places (name_place, type_id) VALUES ($1, $2)`
	_, err = tx.Exec(query, place.NamePlace, typeId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func existsType(tx *sqlx.Tx, typePlace string) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM types_places WHERE type_place = $1)`
	var exists bool
	err := tx.QueryRow(query, typePlace).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("%s: %w", query, err)
	}
	return exists, nil
}

func insertType(tx *sqlx.Tx, place models.Place) error {
	query := `INSERT INTO types_places (type_place) VALUES ($1)`
	_, err := tx.Exec(query, place.Type.Type)
	if err != nil {
		return fmt.Errorf("%s: %w", query, err)
	}
	return nil
}

func (s *Place) GetAll() ([]models.Place, error) {
	const op = "storage.place.GetAll"
	var places []models.Place
	query := `
	SELECT 
    	places.id AS id, 
    	places.name_place AS name_place, 
    	types_places.id AS "type.type_id", 
    	types_places.type_place AS "type.type_place"
	FROM 
	    places
	JOIN 
		types_places 
	ON 
	    types_places.id = places.type_id`

	err := s.db.Select(&places, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return places, nil
}
