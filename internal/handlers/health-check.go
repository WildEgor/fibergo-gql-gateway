package handlers

import (
	"github.com/WildEgor/fibergo-gql-gateway/internal/config"
	domains "github.com/WildEgor/fibergo-gql-gateway/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct {
	appConfig *config.AppConfig
}

func NewHealthCheckHandler(
	appConfig *config.AppConfig,
) *HealthCheckHandler {
	return &HealthCheckHandler{
		appConfig,
	}
}

func (hch *HealthCheckHandler) Handle(c *fiber.Ctx) error {
	c.JSON(fiber.Map{
		"isOk": true,
		"data": &domains.StatusDomain{
			Status: "ok",
		},
	})
	return nil
}
