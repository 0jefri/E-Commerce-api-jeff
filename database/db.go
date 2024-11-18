package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", Cfg.Database.Host, Cfg.Database.Username, Cfg.Database.Password, Cfg.Database.Dbname, Cfg.Database.Port)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Println("errornya:", err)
		log.Fatalf("Error opening database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	// Test koneksi ke database
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// DB = db
	fmt.Println("Successfully Connected To Database!")
}
