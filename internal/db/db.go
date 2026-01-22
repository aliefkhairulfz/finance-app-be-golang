package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Connection struct{}

func (c *Connection) Connect() (*sqlx.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	connectStr := os.Getenv("DB_CONNECTION_STRING")
	db, err := sqlx.Connect("postgres", connectStr)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connection to database successful")
	return db, nil
}
