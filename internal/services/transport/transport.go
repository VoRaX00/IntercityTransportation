package transport

import (
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
	"log/slog"
)

type Repo interface{}

type Transport struct {
	log  *slog.Logger
	repo Repo
}

func New(log *slog.Logger, repo Repo) *Transport {
	return &Transport{
		log:  log,
		repo: repo,
	}
}

func (s *Transport) Add(transport services.AddTransport) error {
	panic("implement me")
}

func (s *Transport) Update(transport services.AddTransport) error {
	panic("implement me")
}

func (s *Transport) Delete(stateNumber string) error {
	panic("implement me")
}

func (s *Transport) GetAll() []models.Transport {
	panic("implement me")
}
