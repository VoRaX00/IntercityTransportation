package place

import (
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
	panic("implement me")
}

func (s *Place) Delete(id int) error {
	panic("implement me")
}

func (s *Place) GetAll() ([]models.Place, error) {
	panic("implement me")
}
