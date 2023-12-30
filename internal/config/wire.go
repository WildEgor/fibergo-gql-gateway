package config

import "github.com/google/wire"

var ConfigsSet = wire.NewSet(
	NewConfigurator,
	NewAppConfig,
	NewGQLConfig,
)
