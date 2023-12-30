package handlers

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/WildEgor/fibergo-gql-gateway/internal/config"
	"github.com/nautilus/gateway"
	"net/http"
)

type GraphQLHandler struct {
	gqlConfig *config.GQLConfig
	gw        *gateway.Gateway
}

func NewGraphQLHandler(gqlConfig *config.GQLConfig) *GraphQLHandler {
	return &GraphQLHandler{
		gqlConfig: gqlConfig,
	}
}

func (g *GraphQLHandler) Init(gw *gateway.Gateway) {
	g.gw = gw
}

func (g *GraphQLHandler) Playground() http.HandlerFunc {
	return g.gw.PlaygroundHandler
}

func (g *GraphQLHandler) Altair() http.HandlerFunc {
	return playground.AltairHandler("GraphQL", g.gqlConfig.Endpoint)
}

func (g *GraphQLHandler) Service() http.HandlerFunc {
	return g.gw.GraphQLHandler
}
