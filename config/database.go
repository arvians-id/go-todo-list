package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

func InitMySQL(configuration Config) (*sql.DB, error) {
	conf := url.Values{}
	conf.Add("parseTime", "True")

	username := configuration.Get("MYSQL_USER")
	password := configuration.Get("MYSQL_PASSWORD")
	host := configuration.Get("MYSQL_HOST")
	port := configuration.Get("MYSQL_PORT")
	database := configuration.Get("MYSQL_DBNAME")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", username, password, host, port, database, conf.Encode())
	db, err := sql.Open(configuration.Get("DB_CONNECTION"), dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
