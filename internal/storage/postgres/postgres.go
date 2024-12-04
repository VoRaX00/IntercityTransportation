package postgres

import "github.com/jmoiron/sqlx"

type Storage struct {
	db *sqlx.DB
}

func New(storagePath string) (*Storage, error) {
	db, err := sqlx.Open("postgres", storagePath)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}
