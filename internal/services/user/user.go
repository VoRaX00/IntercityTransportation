package user

import (
	"fmt"
	"kursachDB/internal/domain/models"
	"log/slog"
)

type Repo interface {
	GetAll() ([]models.User, error)
}

type User struct {
	log  *slog.Logger
	repo Repo
}

func New(log *slog.Logger, repo Repo) *User {
	return &User{
		log:  log,
		repo: repo,
	}
}

func (u *User) GetAll() ([]models.User, error) {
	const op = "user.GetAll"
	log := u.log.With(
		slog.String("op", op),
	)

	log.Info("getting all users")
	users, err := u.repo.GetAll()
	log.Info("returning all users")
	if err != nil {
		log.Warn("error getting all users")
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}
