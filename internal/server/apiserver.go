package server

import (
	"auth/internal/config"
	"net/http"
)

func Start(config *config.Config) error {
	//make conn to db
	// srv := newServer()

	return http.ListenAndServe(config.Port, nil)
}
