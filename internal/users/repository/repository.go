package repository

import (
	"database/sql"
	"finance-app-be/internal/users/model"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

type RepositoryImpl interface {
	FindAll() ([]model.User, error)
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]model.User, error) {
	rows, err := r.db.Query("SELECT id, email, password, is_active FROM users")
	if err != nil {
		fmt.Println("Query failed:", err)
		return nil, err
	}

	defer rows.Close()
	fmt.Println("Database connected and query executed successfully.")

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.IsActive); err != nil {
			fmt.Println("Row scan failed:", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
