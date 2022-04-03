package config

type Config struct {
	Port string `yaml:"port"`
}

func GetConfig() *Config {
	return &Config{}

}
