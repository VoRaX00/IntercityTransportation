package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpUsers, DownUsers)
}

func UpUsers(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS users (
    	phone_number INTEGER PRIMARY KEY,
    	fio TEXT NOT NULL
	);`
	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownUsers(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS users;`
	_, err := tx.ExecContext(ctx, query)
	return err
}
