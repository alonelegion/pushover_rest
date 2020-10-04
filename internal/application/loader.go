package application

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Database initialization
func LoadDB(dialect, url string) (*gorm.DB, error) {
	return gorm.Open(dialect, url)
}

// Load .env file from root directory
func LoadEnv() error {
	return godotenv.Load()
}

// Logrus logger initialization
func InitLogger(level string) *logrus.Logger {
	logger := logrus.New()

	loglevel, err := logrus.ParseLevel(level)
	if err != nil {
		panic("Error setting logger level")
	}

	logger.SetLevel(loglevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	return logger
}
