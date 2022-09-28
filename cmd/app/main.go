package main

import (
	"fmt"
	"log"
	"net/http"

	dbConn "github.com/nathanielhall/cloud-native-go/adapter/db"
	"github.com/nathanielhall/cloud-native-go/app/app"
	"github.com/nathanielhall/cloud-native-go/app/router"
	"github.com/nathanielhall/cloud-native-go/config"
	lr "github.com/nathanielhall/cloud-native-go/util/logger"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	db, err := dbConn.New(appConf)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
		return
	}
	// FIXME: fails since database hasn't been created. Need to setup migrations?
	if err = db.Ping(); err != nil {
		logger.Fatal().Err(err).Msg("")
	} else {
		logger.Debug().Msg("Database ready to accept connections")
	}

	// if appConf.Debug {
	// 	db.LogMode(true)
	// }

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
