package main

import (
	"github.com/alonelegion/pushover_rest/internal/application"
	"github.com/alonelegion/pushover_rest/internal/router"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	// load os.env settings
	if err := application.LoadEnv(); err != nil {
		logrus.Warning(".env file not find")
	}

	// init logrus instance
	logger := application.InitLogger("debug")

	// init gorm instance
	dbURL := os.Getenv("DATABASE_URL")
	logger.Info(os.Getenv("DATABASE_URL"))

	db, err := application.LoadDB(
		os.Getenv("DATABASE_DRIVER"), dbURL+"?sslmode=disable",
	)
	if err != nil {
		logger.Panic(err)
	}

	// init main application singleton instance
	application.Init(db, logger)

	routerEngine := router.NewRouter(logger)
	serveErr := application.Serve(routerEngine, ":"+os.Getenv("APP_PORT"))
	if serveErr != nil {
		application.Logger().Panic(serveErr)
	}

	defer func() {
		if err := db.Close(); err != nil {
			application.Logger().Error(err)
		}
	}()
}
