package app

import (
	"database/sql"

	"github.com/nathanielhall/cloud-native-go/util/logger"
)

type App struct {
	logger *logger.Logger
	db     *sql.DB
}

func New(logger *logger.Logger, db *sql.DB) *App {
	return &App{logger: logger, db: db}
}

func (app *App) Logger() *logger.Logger {
	return app.logger
}
