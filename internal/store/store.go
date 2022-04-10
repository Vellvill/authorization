package store

import (
	"auth/internal/utils"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"
	_ "github.com/lib/pq"
	"time"
)

type Store struct {
	Config         *Config
	Pool           *pgxpool.Pool
	UserRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		Config: config,
		Pool:   nil,
	}
}

// NewClient ...
func (s *Store) NewClient(ctx context.Context) (err error) { //Пул req'ов, подключение возможно не сразу
	dsn := s.Config.DSN
	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		s.Pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil

	}, 5, 5*time.Second)

	if err = s.Pool.Ping(ctx); err != nil {
		return fmt.Errorf("Can't ping db")
	}

	return nil
}

// MigrateDatabase ...
func (s *Store) MigrateDatabase(conn *pgx.Conn, migrationsPath string) (int32, error) {

	migrator, err := migrate.NewMigrator(context.Background(), conn, "schema_version")
	if err != nil {
		return 0, err
	}

	err = migrator.LoadMigrations(migrationsPath)
	if err != nil {
		return 0, err
	}

	err = migrator.Migrate(context.Background())
	if err != nil {
		return 0, err
	}

	ver, err := migrator.GetCurrentVersion(context.Background())
	if err != nil {
		return 0, err
	}

	return ver, nil
}

func (s *Store) Close() {}

// User ...
func (s *Store) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		store: s,
	}

	return s.UserRepository
}

func (s *Store) CLose() {}
