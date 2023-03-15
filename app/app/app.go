package app

import (
	"github.com/nathanielhall/cloud-native-go/util/logger"
	"gorm.io/gorm"
)

type App struct {
	logger *logger.Logger
	db     *gorm.DB
}

const (
    appErrDataAccessFailure   = "data access failure"
    appErrJsonCreationFailure = "json creation failure"
)

func New(logger *logger.Logger, db *gorm.DB) *App {
	return &App{logger: logger, db: db}
}

func (app *App) Logger() *logger.Logger {
	return app.logger
}
