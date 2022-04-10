package store

type Config struct {
	DSN            string `yaml:"dsn"`
	MigrationsPath string `yaml:"migrations_path"`
}

func NewConfig() *Config {
	return &Config{}
}
