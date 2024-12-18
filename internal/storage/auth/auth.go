package auth

import (
	"github.com/jmoiron/sqlx"
	"kursachDB/internal/services"
)

type Auth struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Auth {
	return &Auth{
		db: db,
	}
}

func (s *Auth) AddUser(user services.UserRegister) error {
	panic("implement me")
}
