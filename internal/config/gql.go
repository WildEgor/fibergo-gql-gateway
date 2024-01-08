package config

import (
	"github.com/caarlos0/env/v7"
	log "github.com/sirupsen/logrus"
)

type GQLConfig struct {
	Endpoint           string   `env:"GRAPHQL_ENDPOINT"`
	PlaygroundEndpoint string   `env:"GRAPHQL_PLAYGROUND_ENDPOINT"`
	AltairEndpoint     string   `env:"GRAPHQL_ALTAIR_ENDPOINT"`
	GraphQLEndpoints   []string `env:"GRAPHQL_ENDPOINTS"`
}

func NewGQLConfig(c *Configurator) *GQLConfig {
	cfg := GQLConfig{}

	if err := env.Parse(&cfg); err != nil {
		log.Printf("[GQLConfig] %+v\n", err)
	}

	return &cfg
}
