package main

import (
	"github.com/alonelegion/pushover_rest/internal/application"
	"github.com/sirupsen/logrus"
)

func main() {
	// load os.env settings
	if err := application.LoadEnv(); err != nil {
		logrus.Warning(".env file not find")
	}

	// init main application singleton instance
	app := application.Init()

	InitLogger(app)
	InitDB(app)

	// if something goes wrong, catch it
	defer func() {
		if r := recover(); r != nil {
			app.Logger().
				WithField("RecoverMessage", r).
				Fatal("Application Crashed")
		}
	}()

	InitWebServer(app)
}
