package config

import (
	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
)

type AppConfig struct {
	Name         string `env:"APP_NAME"`
	Port         string `env:"APP_PORT"`
	Mode         string `env:"APP_MODE"`
	Prefork      bool   `env:"PREFORK"`
	ReadTimeout  uint8  `env:"READ_TIMEOUT"`
	ServerHeader string `env:"SERVER_HEADER"`
}

func NewAppConfig(c *Configurator) *AppConfig {
	cfg := AppConfig{}

	if err := env.Parse(&cfg); err != nil {
		log.Printf("[AppConfig] %+v\n", err)
	}

	return &cfg
}

func (ac AppConfig) IsProduction() bool {
	if ac.Mode == "develop" {
		return false
	}

	return true
}
