package server

import "auth/internal/store"

type Config struct {
	Port     string `default:":8080" yaml:"port"`
	LogLevel string `default:"debug" yaml:"log_level"`
	Store    *store.Config
}

func GetConfig() *Config {
	return &Config{
		Port:     ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
