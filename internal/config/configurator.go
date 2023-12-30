package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Configurator struct{}

func NewConfigurator() *Configurator {
	c := &Configurator{}
	if err := c.Load(); err != nil {
		log.Panic("[Configurator] parse error", err)
	}
	return c
}

func (c *Configurator) Load() error {
	err := godotenv.Load(".env", ".env.local")
	return err
}
