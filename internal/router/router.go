package router

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/WildEgor/fibergo-gql-gateway/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/nautilus/gateway"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
)

type GraphqlRouter interface {
	Playground() http.HandlerFunc
	Altair() http.HandlerFunc
	Service() http.HandlerFunc
}

type graphQLHandler struct {
	gw *gateway.Gateway
}

func NewGraphQLHandler(gatewayInst *gateway.Gateway) GraphqlRouter {
	return &graphQLHandler{gw: gatewayInst}
}

func (g *graphQLHandler) Playground() http.HandlerFunc {
	return g.gw.PlaygroundHandler
}

func (g *graphQLHandler) Altair() http.HandlerFunc {
	return playground.AltairHandler("GraphQL", "/api/graphql")
}

func (g *graphQLHandler) Service() http.HandlerFunc {
	return g.gw.GraphQLHandler
}

type Router struct {
	hc *handlers.HealthCheckHandler
}

func NewRouter(hc *handlers.HealthCheckHandler) *Router {
	return &Router{
		hc: hc,
	}
}

func wrapHandler(f func(http.ResponseWriter, *http.Request)) func(ctx *fiber.Ctx) {
	return func(ctx *fiber.Ctx) {
		fasthttpadaptor.NewFastHTTPHandler(http.HandlerFunc(f))(ctx.Context())
	}
}

func (r *Router) Setup(app *fiber.App, gateway *gateway.Gateway) error {
	v1 := app.Group("/api/v1")
	api := app.Group("/api")

	v1.Get("/health", r.hc.Handle)

	gwHandler := NewGraphQLHandler(gateway)

	// Choose one of the following playgrounds
	api.All("/playground", func(c *fiber.Ctx) error {
		wrapHandler(gwHandler.Playground())(c)
		return nil
	})

	// ... or altair (need fix)
	api.All("/altair", func(c *fiber.Ctx) error {
		wrapHandler(gwHandler.Altair())(c)
		return nil
	})

	api.Post("/graphql", func(c *fiber.Ctx) error {
		wrapHandler(gwHandler.Service())(c)
		return nil
	})

	return nil
}
