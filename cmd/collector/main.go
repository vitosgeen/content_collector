package main

import (
	"content_collector/internal/apperrors"
	"content_collector/internal/config"
	"content_collector/internal/infrastructure/datastore"
	"content_collector/internal/infrastructure/logger"
	"content_collector/internal/infrastructure/router"
	"content_collector/internal/interface/controller"
	"content_collector/internal/repository"
	"content_collector/internal/services"
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
)

const (
	env = "./configs/.env"
)

func main() {
	logger := logger.NewLogger()
	cfg, err := config.NewConfig(env)
	if err != nil {
		logger.Fatal(err)
	}

	db, err := datastore.NewClientMongoDB(cfg.Mongo.MongoDbUri, cfg.Mongo.MongoDbUser, cfg.Mongo.MongoDbPass, logger)
	if err != nil {
		logger.Fatal(err)
	}

	collectorRepo := repository.NewCollectorMongoDBRepository(cfg, db, context.TODO())

	fmt.Printf("cfg: %+v\n", cfg)

	collectorService := services.NewCollectorService(cfg.Selenium.ChromeDriverPath, cfg.Selenium.ChromeDriverPort, collectorRepo)
	err = collectorService.CheckCollector()
	if err != nil {
		logger.Fatal(apperrors.ServerStartError.AppendMessage(err))
	}

	controller := controller.NewCollectorController(collectorService)

	e := echo.New()
	e = router.NewRouter(e, controller)
	err = e.Start(cfg.Port)
	if err != nil {
		logger.Fatal(apperrors.ServerStartError.AppendMessage(err))
	}
}
