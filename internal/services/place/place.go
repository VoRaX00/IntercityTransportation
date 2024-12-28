package place

import (
	"errors"
	"fmt"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
	"kursachDB/pkg/mapper"
	"log/slog"
)

var (
	ErrPlaceNotFound = errors.New("place not found")
)

type Repo interface {
	Add(place models.Place) error
	Delete(id int) error
	GetAll() ([]models.Place, error)
}

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
	const op = "Place.Add"
	log := s.log.With(
		slog.String("op", op),
	)

	pl := mapper.PlaceAddToPlace(place)

	log.Info("starting adding place")
	err := s.repo.Add(pl)
	if err != nil {
		log.Warn("failed to add place")
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("place added")
	return nil
}

func (s *Place) Delete(id int) error {
	const op = "Place.Delete"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("starting deleting place")
	err := s.repo.Delete(id)
	if err != nil {
		log.Warn("failed to delete place")
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("place deleted")

	return nil
}

func (s *Place) GetAll() ([]models.Place, error) {
	const op = "Place.GetAll"
	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("starting getting all places")
	places, err := s.repo.GetAll()
	if err != nil {
		log.Warn("failed to get all places")
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	log.Info("successfully got all places")
	return places, nil
}
