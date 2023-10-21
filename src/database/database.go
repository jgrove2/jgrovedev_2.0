package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jgrove2/jgrovedev_2.0/src/util"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDatabase() error {
	var err error
	username := util.GoDotEnvVariable("USER_NAME")
	password := util.GoDotEnvVariable("PASSWORD")
	dsn := fmt.Sprintf("postgresql://%s:%s@jgrove-dev-5905.g8z.cockroachlabs.cloud:26257/jgrove?sslmode=verify-full", username, password)
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	return nil
}

func CloseDatabase() error {
	return DB.Close()
}
