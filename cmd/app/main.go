package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nathanielhall/cloud-native-go/app/app"
	"github.com/nathanielhall/cloud-native-go/app/router"
	"github.com/nathanielhall/cloud-native-go/config"
	lr "github.com/nathanielhall/cloud-native-go/util/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	var logLevel gormlogger.LogLevel

	if appConf.Db.Debug {
		logLevel = gormlogger.Info
	} else {
		logLevel = gormlogger.Error
	}

	dbString := fmt.Sprintf(fmtDBString, appConf.Db.Host, appConf.Db.Username, appConf.Db.Password, appConf.Db.DbName, appConf.Db.Port)
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: gormlogger.Default.LogMode(logLevel)})
	if err != nil {
		logger.Fatal().Err(err).Msg("DB connection start failure")
		return
	} 

	application := app.New(logger, db)

	appRouter := router.New(application)

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
