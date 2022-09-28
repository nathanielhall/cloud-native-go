package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/nathanielhall/cloud-native-go/config"
)

func New(conf *config.Conf) (*sql.DB, error) {

	connInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Db.Host, conf.Db.Username, conf.Db.Password, conf.Db.DbName)

	return sql.Open("pgx", connInfo)
}
