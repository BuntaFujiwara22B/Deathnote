package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("❌ Error conectando a PostgreSQL: %v", err))
	}
	return db
}

func InitDB(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS victimas (
      id SERIAL PRIMARY KEY,
      full_name TEXT NOT NULL,
      cause TEXT,
      details TEXT,
      created_at TIMESTAMP NOT NULL DEFAULT NOW(),
      death_time TIMESTAMP,
      image_url TEXT NOT NULL,
      cause_added BOOLEAN DEFAULT FALSE,
      details_added BOOLEAN DEFAULT FALSE
    )
  `)
	return err
}
