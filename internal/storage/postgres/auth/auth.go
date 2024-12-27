package auth

import (
	"github.com/jmoiron/sqlx"
	"kursachDB/internal/domain/models"
)

type Auth struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func (s *Auth) Login(user models.User) error {
	panic("implement me")
}
