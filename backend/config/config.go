package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	// 🔁 Leer configuración desde variables de entorno (definidas en docker-compose)
	host := os.Getenv("DB_HOST")         // debe ser "postgres" (nombre del servicio en docker-compose)
	port := os.Getenv("DB_PORT")         // "5432"
	user := os.Getenv("DB_USER")         // "postgres"
	password := os.Getenv("DB_PASSWORD") // "bunta"
	dbname := os.Getenv("DB_NAME")       // "deathnote_db"

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

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
      details_added BOOLEAN DEFAULT FALSE,
      is_dead BOOLEAN DEFAULT FALSE
    )
  `)
	return err
}
