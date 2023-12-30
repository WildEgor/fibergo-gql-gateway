package router

import (
	"github.com/WildEgor/fibergo-gql-gateway/internal/handlers"
	"github.com/google/wire"
)

var RouterSet = wire.NewSet(NewRouter, handlers.HandlersSet)
