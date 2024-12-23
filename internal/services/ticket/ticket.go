package ticket

import (
	"kursachDB/internal/domain/models"
	"log/slog"
)

type Repo interface {
}

type Ticket struct {
	log  *slog.Logger
	repo Repo
}

func New(log *slog.Logger, repo Repo) *Ticket {
	return &Ticket{
		log:  log,
		repo: repo,
	}
}

func (t *Ticket) Create() error {
	return nil
}

func (t *Ticket) Update() error {
	return nil
}

func (t *Ticket) Delete() {
}

func (t *Ticket) GetAll() ([]models.Ticket, error) {
	return nil, nil
}
func (t *Ticket) GetByUser() ([]models.Ticket, error) {
	return nil, nil
}
