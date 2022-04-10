package store

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

// TestStore ...
func TestStore(t *testing.T, databaseURL, migrations string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DSN = databaseURL
	config.MigrationsPath = migrations
	s := New(config)
	if err := s.NewClient(context.Background()); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			q := fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
			if _, err := s.Pool.Exec(context.Background(), q); err != nil {
				t.Fatal(err)
			}
		}
		s.Pool.Close()
	}
}
