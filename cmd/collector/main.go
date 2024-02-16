package main

import (
	"context"

	"content_collector/internal/apperrors"
	"content_collector/internal/config"
	"content_collector/internal/infrastructure/datastore"
	"content_collector/internal/infrastructure/logger"
	"content_collector/internal/infrastructure/router"
	"content_collector/internal/interface/controller"
	"content_collector/internal/repository"
	"content_collector/internal/services"
	httpscrapper "content_collector/internal/utils/scrappers/http_scrapper"
	"content_collector/internal/utils/smartproxy"

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

	smartproxyObject := smartproxy.NewSmartProxy(cfg.SmartProxyPath)
	_, err = smartproxyObject.ParseFile()
	if err != nil {
		logger.Fatal(err)
	}

	db, err := datastore.NewClientMongoDB(cfg.Mongo.MongoDbUri, cfg.Mongo.MongoDbUser, cfg.Mongo.MongoDbPass, logger)
	if err != nil {
		logger.Fatal(err)
	}

	scrapper := httpscrapper.NewHttpScpaper()
	// scrapper := seleniumchromescrapper.NewSeleniumChromeScrapper(cfg.Selenium.ChromeDriverPath, cfg.Selenium.ChromeDriverPort)

	collectorRepo := repository.NewCollectorMongoDBRepository(cfg, db, context.TODO())
	collectorService := services.NewCollectorService(smartproxyObject, cfg.Selenium.ChromeDriverPath, cfg.Selenium.ChromeDriverPort, collectorRepo, scrapper)
	controller := controller.NewCollectorController(collectorService)

	e := echo.New()
	router.NewRouter(e, controller)

	err = e.Start(cfg.Port)
	if err != nil {
		logger.Fatal(apperrors.ServerStartError.AppendMessage(err))
	}
}
