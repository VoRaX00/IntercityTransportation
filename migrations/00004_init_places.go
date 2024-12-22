package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpPlaces, DownPlaces)
}

func UpPlaces(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS places(
    	id SERIAL PRIMARY KEY,
    	name_place TEXT NOT NULL,
    	type_id INTEGER REFERENCES types_places(id) NOT NULL
	);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownPlaces(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS places;`
	_, err := tx.ExecContext(ctx, query)
	return err
}
