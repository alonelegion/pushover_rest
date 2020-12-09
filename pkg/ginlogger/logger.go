package ginlogger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var timeFormat = "02/Jan/2006:15:04:05 -0700"

func Logger(logger logrus.FieldLogger) gin.HandlerFunc {
	//hostname := getHostname()

	return func(context *gin.Context) {

	}
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return hostname
}
