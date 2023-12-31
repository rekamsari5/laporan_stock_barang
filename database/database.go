package database

import (
	"database/sql"
	"fmt"
	"laporan_stock_barang/configs"

	"time"

	_ "github.com/go-sql-driver/mysql"
)

func loadDbConfig() (string, error) {

	config := configs.Configs
	if config.DB_HOST == "" {
		return "", fmt.Errorf("Environment variable DB_HOST must be set")
	}
	if config.DB_PORT == "" {
		return "", fmt.Errorf("Environment variable DB_PORT must be set")
	}
	if config.DB_USERNAME == "" {
		return "", fmt.Errorf("Environment variable DB_USERNAME must be set")
	}
	if config.DB_DATABASE == "" {
		return "", fmt.Errorf("Environment variable DB_DATABASE must be set")
	}
	if config.APP_ENV != "local" {
		if config.DB_PASSWORD == "" {
			return "", fmt.Errorf("Environment variable DB_PASSWORD must be set")
		}
	}
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_DATABASE,
	)

	return connStr, nil
}

func GetConnection() (*sql.DB, error) {
	connString, err := loadDbConfig()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db, nil
}
