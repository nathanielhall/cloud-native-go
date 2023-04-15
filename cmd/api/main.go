package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathanielhall/cloud-native-go/api/router"
	"github.com/nathanielhall/cloud-native-go/config"
	"github.com/nathanielhall/cloud-native-go/util/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func main() {
	c := config.New()
	l := logger.New(c.Debug)

	var logLevel gormlogger.LogLevel
	if c.Db.Debug {
		logLevel = gormlogger.Info
	} else {
		logLevel = gormlogger.Error
	}

	dbString := fmt.Sprintf(fmtDBString, c.Db.Host, c.Db.Username, c.Db.Password, c.Db.DbName, c.Db.Port)
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: gormlogger.Default.LogMode(logLevel)})
	if err != nil {
		l.Fatal().Err(err).Msg("DB connection start failure")
		return
	} 

	router := router.New(l, db)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      router,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	closed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		l.Info().Msgf("Shutting down server %v", s.Addr)

		ctx, cancel := context.WithTimeout(context.Background(), c.Server.TimeoutIdle)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			l.Error().Err(err).Msg("Server shutdown failure")
		}

		sqlDB, err := db.DB()
		if err == nil {
			if err = sqlDB.Close(); err != nil {
				l.Error().Err(err).Msg("DB connection closing failure")
			}
		}

		close(closed)
	}()

	l.Info().Msgf("Starting server %v", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		l.Fatal().Err(err).Msg("Server startup failure")
	}

	<-closed
}
