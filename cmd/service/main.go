package main

import (
	"context"
	"log"

	"github.com/pmoura-dev/esr-service/internal/broker"
	"github.com/pmoura-dev/esr-service/internal/config"
	"github.com/pmoura-dev/esr-service/internal/datastore/databases"
	"github.com/pmoura-dev/esr-service/internal/handlers/http_handlers"
	entities_handlers "github.com/pmoura-dev/esr-service/internal/handlers/http_handlers/entities"
	"github.com/pmoura-dev/esr-service/internal/services"
	"github.com/pmoura-dev/esr-service/internal/services/entity"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/gin-gonic/gin"
)

func setupHTTPRouter(entityService services.EntityService) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		http_handlers.EntityService = entityService

		entityGroup := v1.Group("/entities")
		{
			entityGroup.GET("/:entity_id", entities_handlers.GetEntityByID)
			entityGroup.GET("/", entities_handlers.ListEntities)
			entityGroup.POST("/", entities_handlers.AddEntity)
			entityGroup.DELETE("/:entity_id", entities_handlers.DeleteEntity)
			entityGroup.POST("/:entity_id/commands", entities_handlers.NewCommand)
		}
	}
	return router
}

func setupPubSubRouter() (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, watermill.NewSlogLogger(nil))
	if err != nil {
		return nil, err
	}

	router.AddPlugin(plugin.SignalsHandler)

	return router, nil
}

func main() {

	cfg := config.LoadConfig()

	// Initialize datastore
	db, err := databases.GetDataStore(cfg.DataStore)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	// Initialize broker
	bk, err := broker.GetBroker(cfg.Broker)
	if err != nil {
		log.Fatal(err)
	}
	defer bk.Close()

	// Services
	entityService := entity.NewBaseEntityService(db, bk)

	httpRouter := setupHTTPRouter(entityService)
	go func() {
		if err := httpRouter.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	pubSubRouter, err := setupPubSubRouter()
	if err != nil {
		log.Fatal(err)
	}

	if err := pubSubRouter.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
