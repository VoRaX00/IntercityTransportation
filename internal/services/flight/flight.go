package flight

import (
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
	"log/slog"
)

type Repo interface{}

type Flight struct {
	log  *slog.Logger
	repo Repo
}

func New(log *slog.Logger, repo Repo) *Flight {
	return &Flight{
		log:  log,
		repo: repo,
	}
}

func (s *Flight) Add(flight services.AddFlight) error {
	panic("implement me")
}

func (s *Flight) Delete(id int) error {
	panic("implement me")
}

func (s *Flight) GetAll() ([]models.Flight, error) {
	panic("implement me")
}
