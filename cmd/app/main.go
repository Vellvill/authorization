package main

import (
	config "auth/internal/config"
	"auth/internal/server"
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yaml", "path to config file")
}

func main() {
	flag.Parse()

	conf := config.GetConfig()

	if err := cleanenv.ReadConfig(configPath, conf); err != nil {
		_, err = cleanenv.GetDescription(conf, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := server.Start(conf); err != nil {
		log.Fatal(err)
	}
}
