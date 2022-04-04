package server

type Config struct {
	Port     string `yaml:"port"`
	LogLevel string `yaml:"log_level"`
	Store    struct {
		DSN string `yaml:"dsn"`
	} `yaml:"store"`
}

func GetConfig() *Config {
	return &Config{
		Port:     ":8080",
		LogLevel: "debug",
	}

}
