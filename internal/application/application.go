package application

import (
	"errors"
	"github.com/alonelegion/pushover_rest/internal/queries"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/now"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

type Application struct {
	db           *gorm.DB
	logger       *logrus.Logger
	envMode      EnvironmentMode
	Dependencies *Dependencies
}

type Dependencies struct {
	BaseQuery queries.Query
}

var (
	instance *Application
	once     sync.Once
)

// Application errors
var (
	ErrInvalidEnvMode = errors.New("invalid environment mode")
)

// Initializing a Application
func Init(db *gorm.DB, logger *logrus.Logger) *Application {
	once.Do(func() {
		envMode, errEnv := receiveEnvironmentMode()
		if errEnv != nil {
			logger.Warn(errEnv)
		}

		deps := &Dependencies{
			BaseQuery: queries.InitQuery(db),
		}
		instance = &Application{
			logger:       logger,
			db:           db,
			envMode:      envMode,
			Dependencies: deps,
		}

		// set default week start day to monday
		now.WeekStartDay = time.Monday
	})
	return instance
}

func Serve(r *gin.Engine, port string) error {
	err := r.Run(port)
	if err != nil {
		return err
	}

	return nil
}

// Receive and validate environment mode from .env file
func receiveEnvironmentMode() (EnvironmentMode, error) {
	envMode := EnvironmentMode(os.Getenv("APP_ENV"))

	switch envMode {
	case Development, Production:
		return envMode, nil
	}

	return Unknown, ErrInvalidEnvMode
}

func Get() *Application {
	return instance
}

func (a *Application) Db() *gorm.DB {
	return a.db
}

func (a *Application) Logger() *logrus.Logger {
	return a.logger
}

func Logger() *logrus.Logger {
	return Get().Logger()
}

func DB() *gorm.DB {
	return Get().Db()
}
