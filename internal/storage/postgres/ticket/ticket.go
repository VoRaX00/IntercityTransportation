package ticket

import (
	"github.com/jmoiron/sqlx"
	"kursachDB/internal/domain/models"
)

type Ticket struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Ticket {
	return &Ticket{
		db: db,
	}
}

func (r *Ticket) Add(ticket models.Ticket) error {
	panic("implement me")
}

func (r *Ticket) Delete(phoneNumber int64) error {
	panic("implement me")
}

func (r *Ticket) GetAll() ([]models.Ticket, error) {
	panic("implement me")
}

func (r *Ticket) GetByUser(phoneNumber int64) ([]models.Ticket, error) {
	panic("implement me")
}
