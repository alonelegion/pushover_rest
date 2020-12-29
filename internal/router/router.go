package router

import (
	v1 "github.com/alonelegion/pushover_rest/internal/controllers/v1"
	"github.com/alonelegion/pushover_rest/pkg/gincache"
	"github.com/alonelegion/pushover_rest/pkg/ginlogger"
	"github.com/alonelegion/pushover_rest/pkg/recoverywriter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

func NewRouter(logger *logrus.Logger) *gin.Engine {
	router := gin.New()
	router.Use(ginlogger.Logger(logger))

	pc := new(v1.PingController)
	router.GET("/", pc.Ping)

	if os.Getenv("SLEEPER") == "true" {
		router.Use(Sleeper())
	}

	// Api
	apiRouter := router.Group("api")
	apiRouter.Use(gin.RecoveryWithWriter(recoverywriter.NewGinRecoverWriter(logger)))
	apiRouter.Use(gincache.CacheResponse(logger))
	mapV1Routes(apiRouter)

	return router
}

func mapV1Routes(router *gin.RouterGroup) {
	v1Group := router.Group("v1")
	{
		// Pushover
		pushoverGroup := v1Group.Group("pushover")
		{
			testing := new(v1.PushoverController)
			pushoverGroup.GET("/check", testing.Check)
		}
	}
}
