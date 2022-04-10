package store_test

import (
	"auth/internal/model"
	"auth/internal/store"
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, DSN, MigrationsPath)
	defer teardown("users")

	conn, err := s.Pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Unable to acquire a test database connection: %v\n", err)
	}

	_, err = s.MigrateDatabase(conn.Conn(), MigrationsPath)
	if err != nil {
		log.Fatalf("Unable to migrate a test database connection: %v\n", err)
	}

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, DSN, MigrationsPath)
	defer teardown("users")

	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(&model.User{
		Email: "user@example.org",
	})

	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
