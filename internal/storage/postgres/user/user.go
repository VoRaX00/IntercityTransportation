package user

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"kursachDB/internal/domain/models"
)

type User struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) Login(user models.User) error {
	const op = "postgres.user.Login"
	_, err := u.Get(user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			query := `INSERT INTO user (phone_number, fio) VALUES ($1, $2)`
			_, err = u.db.Query(query, user.PhoneNumber, user.FIO)
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
			return nil
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (u *User) Get(user models.User) (models.User, error) {
	const op = "postgres.user.Get"
	query := `SELECT phone_number, fio FROM user WHERE phone_number = $1`
	var res models.User
	err := u.db.Get(&res, query, user.PhoneNumber)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	if res.FIO != user.FIO {
		return models.User{}, fmt.Errorf("%s: %w", op, fmt.Errorf("%s: %w", op, "FIO does not match"))
	}
	return res, nil
}

func (u *User) GetAll() ([]models.User, error) {
	const op = "postgres.user.GetAll"
	query := `SELECT * FROM users`

	var users []models.User
	err := u.db.Select(&users, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return users, nil
}
