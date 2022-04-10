package store_test

import (
	"os"
	"testing"
)

var (
	DSN            string
	MigrationsPath string
)

func TestMain(m *testing.M) {
	DSN, MigrationsPath = os.Getenv("DATABASE_URL"), os.Getenv("MIGRATIONS")
	if DSN == "" || MigrationsPath == "" {
		DSN = "postgresql://timofey:12345@localhost:5432/restapi_test"
		MigrationsPath = "./migrations_test"
	}

	os.Exit(m.Run())
}
