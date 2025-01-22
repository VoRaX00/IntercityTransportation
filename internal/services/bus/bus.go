package bus

import (
	"errors"
	"fmt"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
	"log/slog"
)

var (
	ErrBusNotFound = errors.New("bus not found")
)

type Repo interface {
	Add(bus models.Bus) error
	Get(stateNumber string) (models.Bus, error)
	GetAll() ([]models.Bus, error)
}

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
	const op = "bus.Add"
	log := s.log.With(
		slog.String("op", op),
	)

	param := models.Bus{
		StateNumber: bus.StateNumber,
		Model: models.Model{
			Model: bus.Model,
		},
	}

	log.Info("starting add bus")
	err := s.repo.Add(param)
	if err != nil {
		log.Warn("error adding bus", err.Error())
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("success add bus")
	return nil
}

func (s *Bus) Get(stateNumber string) (models.Bus, error) {
	const op = "bus.Get"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("starting get bus")
	bus, err := s.repo.Get(stateNumber)
	if err != nil {
		log.Warn("error getting bus", err)
		return models.Bus{}, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("success get bus")
	return bus, nil
}

func (s *Bus) GetAll() ([]models.Bus, error) {
	const op = "bus.GetAll"
	log := s.log.With(slog.String("op", op))

	log.Info("fetching all bus")
	buses, err := s.repo.GetAll()
	if err != nil {
		log.Warn("error getting buses", err.Error())
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("returning all bus")
	return buses, nil
}
