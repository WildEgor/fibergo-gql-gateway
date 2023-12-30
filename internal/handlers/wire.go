package handlers

import (
	"github.com/google/wire"
)

var HandlersSet = wire.NewSet(NewHealthCheckHandler)
