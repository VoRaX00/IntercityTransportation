package ticket

import (
	"fmt"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
	"kursachDB/pkg/mapper"
	"log/slog"
)

type Repo interface {
	Add(ticket models.Ticket) error
	Delete(phoneNumber int64) error
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

	tick := mapper.BuyTicketToTicket(ticket)

	log.Info("start creating ticket")
	err := s.repo.Add(tick)
	if err != nil {
		log.Warn("failed to add ticket", err)
	}
	log.Info("finish creating ticket")
	return nil
}

func (s *Ticket) RemoveTicket(phoneNumber int64) error {
	const op = "Ticket.Create"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("start deleting ticket")
	err := s.repo.Delete(phoneNumber)
	if err != nil {
		log.Error("failed to delete ticket", err)
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("finish deleting ticket")
	return nil
}

func (s *Ticket) GetAll() ([]models.Ticket, error) {
	return nil, nil
}
func (s *Ticket) GetByUser(phoneNumber int64) ([]models.Ticket, error) {
	return nil, nil
}
