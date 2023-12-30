// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package pkg

import (
	"github.com/WildEgor/fibergo-gql-gateway/internal/config"
	"github.com/WildEgor/fibergo-gql-gateway/internal/handlers"
	"github.com/WildEgor/fibergo-gql-gateway/internal/router"
	"github.com/google/wire"
)

// Injectors from server.go:

func NewServer() (*Server, error) {
	configurator := config.NewConfigurator()
	appConfig := config.NewAppConfig(configurator)
	gqlConfig := config.NewGQLConfig(configurator)
	healthCheckHandler := handlers.NewHealthCheckHandler(appConfig)
	graphQLHandler := handlers.NewGraphQLHandler(gqlConfig)
	routerRouter := router.NewRouter(healthCheckHandler, graphQLHandler, gqlConfig)
	server := NewApp(appConfig, gqlConfig, routerRouter)
	return server, nil
}

// server.go:

var ServerSet = wire.NewSet(AppSet)
