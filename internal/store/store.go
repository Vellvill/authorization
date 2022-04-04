package store

import (
	"context"
	"database/sql"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"

	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DSN)
	if err != nil {
		return nil
	}

	if err = db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
