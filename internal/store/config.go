package store

type Config struct {
	DSN string `yaml:"dsn"`
}

func NewConfig() *Config {
	return &Config{}
}
