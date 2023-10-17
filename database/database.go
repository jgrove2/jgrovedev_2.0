package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDatabase() error {
	var err error
	DB, err = sql.Open("postgres", "user=postgres password=eGwauP8E8VwCYY host=db.dzigthayjnuozieeimvt.supabase.co port=5432 dbname=postgres sslmode=disable")
	log.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func CloseDatabase() error {
	return DB.Close()
}
