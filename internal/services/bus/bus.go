package bus

import (
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
	"log/slog"
)

type Repo interface{}

type Bus struct {
	log  *slog.Logger
	repo Repo
}

func New(log *slog.Logger, repo Repo) *Bus {
	return &Bus{
		log:  log,
		repo: repo,
	}
}

func (s *Bus) Add(bus services.AddBus) error {
	panic("implement me")
}

func (s *Bus) Update(bus services.AddBus) error {
	panic("implement me")
}

func (s *Bus) Delete(stateNumber string) error {
	panic("implement me")
}

func (s *Bus) GetAll() []models.Bus {
	panic("implement me")
}
