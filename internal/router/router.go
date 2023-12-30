package router

import (
	"github.com/WildEgor/fibergo-gql-gateway/internal/config"
	"github.com/WildEgor/fibergo-gql-gateway/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/nautilus/gateway"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
)

type Router struct {
	hc        *handlers.HealthCheckHandler
	gql       *handlers.GraphQLHandler
	gqlConfig *config.GQLConfig
}

func NewRouter(hc *handlers.HealthCheckHandler, gql *handlers.GraphQLHandler, gqlConfig *config.GQLConfig) *Router {
	return &Router{
		hc:        hc,
		gql:       gql,
		gqlConfig: gqlConfig,
	}
}

func wrapHandler(f func(http.ResponseWriter, *http.Request)) func(ctx *fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(f))(ctx.Context())
	}
}

func (r *Router) Setup(app *fiber.App, gateway *gateway.Gateway) error {
	v1 := app.Group("/api/v1")
	api := app.Group("/")

	v1.Get("/health", r.hc.Handle)

	r.gql.Init(gateway)

	// Choose one of the following playgrounds
	api.All(r.gqlConfig.PlaygroundEndpoint, func(c *fiber.Ctx) error {
		wrapHandler(r.gql.Playground())(c)
		return nil
	})

	// ... or altair (need fix)
	api.All(r.gqlConfig.AltairEndpoint, func(c *fiber.Ctx) error {
		wrapHandler(r.gql.Altair())(c)
		return nil
	})

	api.Post(r.gqlConfig.Endpoint, func(c *fiber.Ctx) error {
		wrapHandler(r.gql.Service())(c)
		return nil
	})

	return nil
}
