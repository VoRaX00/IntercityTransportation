package schedule

import (
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
	"log/slog"
)

type Repo interface{}

type Schedule struct {
	log  *slog.Logger
	repo Repo
}

func New(log *slog.Logger, repo Repo) *Schedule {
	return &Schedule{
		log:  log,
		repo: repo,
	}
}

func (s *Schedule) Add(schedule services.AddSchedule) error {
	panic("implement me")
}

func (s *Schedule) Delete(id int) error {
	panic("implement me")
}

func (s *Schedule) GetAll() []models.Schedule {
	panic("implement me")
}
