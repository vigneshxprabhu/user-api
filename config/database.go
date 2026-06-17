package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDB() *sql.DB {

	db, err := sql.Open(
		"postgres",
		"postgres://postgres:yolo123@localhost:5432/user_api?sslmode=disable",
	)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")

	return db
}
