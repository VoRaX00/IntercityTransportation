package auth

import (
	"database/sql"
	"errors"
	"fmt"
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
	const op = "storage.auth.login"
	_, err := s.Get(user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query := `INSERT INTO users (phone_number, fio) VALUES ($1, $2)`
			_, err = s.db.Query(query, user.PhoneNumber, user.FIO)
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
			return nil
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *Auth) Get(user models.User) (models.User, error) {
	const op = "storage.auth.getByPhoneNumber"
	query := `SELECT phone_number, fio FROM users WHERE phone_number = $1`
	var res models.User
	err := s.db.Get(&res, query, user.PhoneNumber)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	if res.FIO != user.FIO {
		return models.User{}, fmt.Errorf("%s: %w", op, fmt.Errorf("%s: %w", op, "FIO does not match"))
	}
	return res, nil
}
