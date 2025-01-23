package auth

import (
	"fmt"
	"kursachDB/internal/domain/models"
	"kursachDB/pkg/jwt"
	"log/slog"
	"time"
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

func (s *Auth) Login(login models.User) (string, error) {
	const op = "signIn"

	log := s.log.With(
		slog.String("op", op),
	)

	log.Info("login verification")
	err := s.repo.Login(login)
	if err != nil {
		log.Warn("error login", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("login verified")

	log.Info("creating jwt")
	token, err := jwt.NewToken(login, time.Hour*24)
	if err != nil {
		log.Warn("error creating jwt", err)
		return "", fmt.Errorf("%s: %w", op, err)
	}
	log.Info("jwt created")

	return token, nil
}
