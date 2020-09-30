package application

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

type Application struct {
	db     *gorm.DB
	logger *logrus.Logger
}

var (
	instance *Application
	once     sync.Once
)

// Initializing a Application
func Init(db *gorm.DB, logger *logrus.Logger) *Application {
	return instance
}

func Serve(r *gin.Engine, port string) error {
	err := r.Run(port)
	if err != nil {
		return err
	}

	return nil
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
