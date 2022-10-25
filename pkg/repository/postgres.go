package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	cityTable = "cityinfo"
	cityTemp  = "citytemp"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", cfg.Username, cfg.Password,
		cfg.DBName, "disable"))

	if err != nil {
		return nil, err
	}

	//Checking db connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
