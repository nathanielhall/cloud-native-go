package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/nathanielhall/cloud-native-go/config"
	"github.com/nathanielhall/cloud-native-go/util/logger"
)

func NewDb(conf *config.Conf, lr *logger.Logger) (*sql.DB, error) {
	connInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Db.Host, conf.Db.Username, conf.Db.Password, conf.Db.DbName)

	db, err := sql.Open("pgx", connInfo)
	if err != nil {
		lr.Fatal().Err(err).Msg("Database connection error")
	}

	return db, err
}
