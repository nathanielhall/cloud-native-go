package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nathanielhall/cloud-native-go/app/router"
	"github.com/nathanielhall/cloud-native-go/config"
)
 

func main() {
	appConf := config.AppConfig()
	appRouter := router.New()

	address := fmt.Sprintf(":%d", appConf.Server.Port)
	log.Printf("Starting server %s\n", address)

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
 