package helpers

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func SetupPostgres() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		GetEnv("DB_HOST", "127.0.0.1"),
		GetEnv("DB_USER", ""),
		GetEnv("DB_PASSWORD", ""),
		GetEnv("DB_NAME", ""),
		GetEnv("DB_PORT", "5432"),
	)

	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	log.Println("Successfully connected to PostgreSQL database...")
}
