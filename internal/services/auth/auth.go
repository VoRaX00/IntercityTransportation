package auth

import (
	"fmt"
	"kursachDB/internal/domain/models"
	"log/slog"
)

type Repo interface {
	Login(user models.User) error
}

type Auth struct {
	log  *slog.Logger
	repo Repo
}

func New(log *slog.Logger, repo Repo) *Auth {
	return &Auth{
		log:  log,
		repo: repo,
	}
}

func (s *Auth) Login(login models.User) error {
	const op = "signIn"

	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("login verification")
	err := s.repo.Login(login)
	if err != nil {
		log.Warn("error login", err)
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("login verified")
	return nil
}
