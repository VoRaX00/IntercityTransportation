package place

import (
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
	"log/slog"
)

type Repo interface{}

type Place struct {
	log  *slog.Logger
	repo Repo
}

func New(log *slog.Logger, repo Repo) *Place {
	return &Place{
		log:  log,
		repo: repo,
	}
}

func (s *Place) Add(place services.AddPlace) error {
	panic("implement me")
}

func (s *Place) Delete(id int) error {
	panic("implement me")
}

func (s *Place) GetAll() []models.Place {
	panic("implement me")
}
