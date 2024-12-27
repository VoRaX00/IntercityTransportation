package ticket

import (
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
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

func (s *Ticket) BuyTicket(ticket services.BuyTicket) error {
	const op = "Ticket.Create"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("start creating ticket")

	return nil
}

func (t *Ticket) Update() error {
	return nil
}

func (t *Ticket) RemoveTicket(id int64) error {
	panic("implement me")
}

func (t *Ticket) GetAll() ([]models.Ticket, error) {
	return nil, nil
}
func (t *Ticket) GetByUser() ([]models.Ticket, error) {
	return nil, nil
}
