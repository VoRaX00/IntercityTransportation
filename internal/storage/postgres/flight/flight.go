package flight

import (
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
	panic("implement me")
}

func (r *Flight) Delete(id int) error {
	panic("implement me")
}

func (r *Flight) GetAll() ([]models.Flight, error) {
	panic("implement me")
}
