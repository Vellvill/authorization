package store

import (
	"auth/internal/utils"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"log"
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
		log.Fatal(err)
	}

	return nil
}

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
