package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nathanielhall/cloud-native-go/app/app"
	"github.com/nathanielhall/cloud-native-go/app/router"
	"github.com/nathanielhall/cloud-native-go/config"
	lr "github.com/nathanielhall/cloud-native-go/util/logger"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	application := app.New(logger)
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
