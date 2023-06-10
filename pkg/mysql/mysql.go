package mysql

import (
	"fmt"
	"log"
	"utils/mod-user-dash/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(cfg *config.Config) (*sqlx.DB, error) {
	log.Println("Trying to connect with MySQL Server")
	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		cfg.MySQL.Username,
		cfg.MySQL.Password,
		cfg.MySQL.Host,
		cfg.MySQL.Port,
		cfg.MySQL.Database,
	)

	db, err := sqlx.Connect(`mysql`, dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, err
}
