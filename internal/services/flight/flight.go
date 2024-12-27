package flight

import (
	"fmt"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
	"kursachDB/pkg/mapper"
	"log/slog"
)

type Repo interface {
	Add(flight models.Flight) error
	Delete(id int) error
	GetAll() ([]models.Flight, error)
}

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
	const op = "Flight.Add"
	log := s.log.With(
		slog.String("op", op),
	)

	fl, err := mapper.FlightAddToFlight(flight)
	if err != nil {
		log.Warn(err.Error())
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("starting adding flight")
	err = s.repo.Add(fl)
	if err != nil {
		log.Warn("error adding flight", err)
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("finishing adding flight")

	return nil
}

func (s *Flight) Delete(id int) error {
	const op = "Flight.Delete"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("starting deleting flight")
	err := s.repo.Delete(id)
	if err != nil {
		log.Warn("error deleting", err)
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("finishing deleting flight")
	return nil
}

func (s *Flight) GetAll() ([]models.Flight, error) {
	const op = "Flight.GetAll"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("starting getting all flights")
	flights, err := s.repo.GetAll()
	if err != nil {
		log.Warn("error getting all flights", err)
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("finishing getting all flights")
	return flights, nil
}
