package main

import (
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

	conf := server.GetConfig()

	if err := cleanenv.ReadConfig(configPath, conf); err != nil {
		_, err = cleanenv.GetDescription(conf, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	s := server.New(conf)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
