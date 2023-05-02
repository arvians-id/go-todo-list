package config

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/url"
	"time"
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

	dbpool, err := databasePooling(db)
	if err != nil {
		return nil, err
	}

	//err = createTable(dbpool)
	//if err != nil {
	//	return nil, err
	//}

	return dbpool, nil
}

func createTable(db *sql.DB) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, `
			CREATE TABLE IF NOT EXISTS activities(
			   activity_id INTEGER NOT NULL AUTO_INCREMENT,
			   title VARCHAR (100) NOT NULL,
			   email VARCHAR (100) UNIQUE NOT NULL,
			   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			   PRIMARY KEY (activity_id)
			) ENGINE = InnoDB;
	`)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)

	_, err = db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS todos(
		   todo_id INTEGER NOT NULL AUTO_INCREMENT,
		   activity_group_id INTEGER NOT NULL,
		   title VARCHAR (100) NOT NULL,
		   priority VARCHAR (100) NOT NULL,
		   is_active BOOLEAN NOT NULL,
		   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		   PRIMARY KEY (todo_id),
		   FOREIGN KEY (activity_group_id) REFERENCES activities(activity_id) ON UPDATE CASCADE ON DELETE CASCADE
		) ENGINE = InnoDB;
	`)
	if err != nil {
		return err
	}
	rows, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)

	return nil
}

func databasePooling(db *sql.DB) (*sql.DB, error) {
	// Limit connection with db pooling
	db.SetMaxIdleConns(20)                 // minimal connection
	db.SetMaxOpenConns(20)                 // maximal connection
	db.SetConnMaxLifetime(5 * time.Minute) // unused connections will be deleted
	db.SetConnMaxIdleTime(1 * time.Minute) // connection that can be used

	return db, nil
}
