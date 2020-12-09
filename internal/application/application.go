package application

import (
	"errors"
	"github.com/alonelegion/pushover_rest/internal/queries"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/now"
	"sync"
	"time"

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
	// Current Application
	instance *Application
	once     sync.Once
)

// Application errors
var (
	ErrInvalidEnvMode = errors.New("invalid environment mode")
)

// Initializing a Application
func Init() *Application {
	once.Do(func() {
		envMode, errEnv := receiveEnvironmentMode()
		if errEnv != nil {
			logrus.Warn(errEnv)
		}

		instance = &Application{
			envMode:      envMode,
			Dependencies: &Dependencies{},
		}

		// set default week start day to monday
		now.WeekStartDay = time.Monday
	})
	return instance
}

func Get() *Application {
	return instance
}

func (a *Application) DB() *gorm.DB {
	return a.db
}

func (a *Application) SetDB(db *gorm.DB) {
	a.db = db
}

func (a *Application) Logger() *logrus.Logger {
	return a.logger
}

func (a *Application) SetLogger(logger *logrus.Logger) {
	a.logger = logger
}

func (a *Application) EnvMode() EnvironmentMode {
	return a.envMode
}

func (a *Application) Deps() *Dependencies {
	return a.Dependencies
}

func Logger() *logrus.Logger {
	return Get().Logger()
}

func DB() *gorm.DB {
	return Get().DB()
}

func EnvMode() EnvironmentMode {
	return Get().EnvMode()
}

func Deps() *Dependencies {
	return Get().Deps()
}
