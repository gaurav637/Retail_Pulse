package database

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kiranaClub/config"
)

// DB is a global variable for database connection
var DB *sqlx.DB

// DBConnecting initializes and returns a new database connection
func DBConnecting(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	log.Println("Database connected successfully")
	DB = db
	return db, nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("failed to close database connection: %v", err)
			return err
		}
		log.Println("Database connection closed")
		DB = nil
	}
	return nil
}
