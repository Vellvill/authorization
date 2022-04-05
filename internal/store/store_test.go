package store

import (
	"os"
	"testing"
)

var (
	DSN string
)

func TestMain(m *testing.M) {
	DSN = os.Getenv("DATABASE_URL")
	if DSN == "" {
		DSN = "postgresql://postgres:12345@localhost:5432/postgres"
	}

	os.Exit(m.Run())
}
