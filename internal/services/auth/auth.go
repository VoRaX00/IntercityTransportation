package auth

import (
	"kursachDB/internal/services"
	"log/slog"
)

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

func (s *Auth) SignIn(login services.UserLogin) error {
	panic("implement me")
}
