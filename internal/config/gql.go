package config

import (
	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
)

type GQLConfig struct {
	GraphQLEndpoints []string `env:"GRAPHQL_ENDPOINTS"`
}

func NewGQLConfig(c *Configurator) *GQLConfig {
	cfg := GQLConfig{}

	if err := c.Load(); err == nil {
		if err := env.Parse(&cfg); err != nil {
			log.Printf("[GQLConfig] %+v\n", err)
		}
	}

	return &cfg
}
