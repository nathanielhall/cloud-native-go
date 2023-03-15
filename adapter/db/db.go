package db

import (
	"fmt"

	"github.com/nathanielhall/cloud-native-go/config"
	"github.com/nathanielhall/cloud-native-go/util/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

func NewDb(conf *config.Conf, lr *logger.Logger) (*gorm.DB, error) {
	var logLevel gormlogger.LogLevel

	if conf.Db.Debug {
		logLevel = gormlogger.Info
	} else {
		logLevel = gormlogger.Error
	}

	dbString := fmt.Sprintf(fmtDBString, conf.Db.Host, conf.Db.Username, conf.Db.Password, conf.Db.DbName, conf.Db.Port)

	lr.Info().Msgf("Database string %v", dbString)

	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: gormlogger.Default.LogMode(logLevel)})
	return db, err
}
