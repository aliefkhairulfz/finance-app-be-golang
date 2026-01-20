package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Connection struct{}

func (c *Connection) Connect() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	connectStr := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connection to database successful")
	return db, nil
}
