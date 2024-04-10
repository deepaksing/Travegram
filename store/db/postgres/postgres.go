package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/deepaksing/Travegram/store"
)

type DB struct {
	db *sql.DB
}

func NewDB() (store.Driver, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/Travegram?sslmode=disable")
	if err != nil {
		return nil, err
	}
	var driver store.Driver = &DB{
		db: db,
	}
	return driver, nil
}

// Migrate executes database migrations
func (d *DB) Migrate(ctx context.Context) error {
	buf, err := os.ReadFile("store/db/postgres/SCHEMA.sql")
	if err != nil {
		return fmt.Errorf("failed to read latest schema file: %w", err)
	}
	stmt := string(buf)
	_, err = d.db.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	return nil
}
