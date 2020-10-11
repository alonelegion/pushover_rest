package router

import (
	v1 "github.com/alonelegion/pushover_rest/internal/controllers/v1"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

func NewRouter(logger *logrus.Logger) *gin.Engine {
	router := gin.New()

	pc := new(v1.PingController)
	router.GET("/", pc.Ping)

	if os.Getenv("SLEEPER") == "true" {
		router.Use(Sleeper())
	}

	return router
}
