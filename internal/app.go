package pkg

import (
	"github.com/WildEgor/fibergo-gql-gateway/internal/config"
	"github.com/WildEgor/fibergo-gql-gateway/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"
	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
	"os"

	log "github.com/sirupsen/logrus"
)

var AppSet = wire.NewSet(
	NewApp,
	config.ConfigsSet,
	router.RouterSet,
)

var Version = "1.0.0"

type Server struct {
	App       *fiber.App
	AppConfig *config.AppConfig
}

func NewApp(
	appConfig *config.AppConfig,
	gqlConfig *config.GQLConfig,
	router *router.Router,
) *Server {
	app := fiber.New(fiber.Config{
		Prefork: appConfig.Prefork,
		// ReadTimeout:  time.Second * time.Duration(appConfig.ReadTimeout),
		AppName:      appConfig.Name + " Version: " + Version,
		ServerHeader: appConfig.ServerHeader,
		// ErrorHandler: error_handler.ErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST",
	}))
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// Set logging settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	if !appConfig.IsProduction() {
		// HINT: some extra setting
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	/// Gateway Init
	// Load the endpoints here first before anything...
	schemas, err := graphql.IntrospectRemoteSchemas(gqlConfig.GraphQLEndpoints...)
	if err != nil {
		log.Fatalf("[Bootstrap] Schema loading error: %s", err.Error())
	}

	log.Println("[Bootstrap] Schema loaded successfully")

	gw, err := gateway.New(schemas)
	if err != nil {
		log.Fatalf("[Bootstrap] Gateway initialization error: %s", err.Error())
	}

	log.Println("[Bootstrap] Gateway initialized successfully")

	err = router.Setup(app, gw)
	if err != nil {
		log.Fatalf("[Bootstrap] Router initialization error: %s", err.Error())
	}

	log.Println("[Bootstrap] Router initialized successfully")

	return &Server{
		App:       app,
		AppConfig: appConfig,
	}
}
