package repository

import (
	"database/sql"
	"errors"
	"finance-app-be/internal/users/model"
	"fmt"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	db *sqlx.DB
}

type RepositoryImpl interface {
	FindAll() ([]model.UserResponse, error)
	FindOneById(id int) (*model.UserResponse, error)
	FindOneByEmail(email string) (*model.UserResponse, error)
	Create(email, password string) (*int64, error)
	UpdateOneById(id int, avatar *string, isActive *bool) (*int64, error)
}

func NewRepository(db *sqlx.DB) RepositoryImpl {
	return &Repository{db: db}
}

func (r *Repository) FindOneById(id int) (*model.UserResponse, error) {
	var user model.UserResponse
	sqlStatement := `SELECT id, email, avatar, is_active, created_at, updated_at FROM users WHERE id=$1`

	if err := r.db.Get(&user, sqlStatement, id); err != nil {
		return nil, err
	}

	fmt.Println("Query executed successfully.")
	return &user, nil
}

func (r *Repository) FindOneByEmail(email string) (*model.UserResponse, error) {
	var user model.UserResponse
	sqlStatement := `SELECT id, email, avatar, is_active, created_at, updated_at FROM users WHERE email=$1`

	if err := r.db.Get(&user, sqlStatement, email); err != nil {
		return nil, err
	}

	fmt.Println("Query executed successfully.")
	return &user, nil
}

func (r *Repository) FindAll() ([]model.UserResponse, error) {
	var users []model.UserResponse
	err := r.db.Select(&users, "SELECT id, email, avatar, is_active, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}

	fmt.Println("Query executed successfully.")
	return users, nil
}

func (r *Repository) Create(email, password string) (*int64, error) {
	findUser, errFind := r.FindOneByEmail(email)
	if errFind != sql.ErrNoRows && findUser != nil {
		fmt.Println("User already exists with email:", email)
		return nil, errors.New("user already exists with email: " + email)
	}

	hshdPasswd, errHshd := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errHshd != nil {
		fmt.Println("Password hashing failed:", errHshd)
		return nil, errHshd
	}

	res, err := r.db.NamedExec("INSERT INTO users (email, password) VALUES (:email, :password)",
		map[string]interface{}{
			"email":    email,
			"password": hshdPasswd,
		})

	if err != nil {
		fmt.Println("Insert failed:", err)
		return nil, err
	}

	fmt.Println("Query executed successfully.")

	affected, errAfct := res.RowsAffected()
	if errAfct != nil {
		fmt.Println("Retrieving affected rows failed:", errAfct)
		return nil, errAfct
	}

	return &affected, nil
}

func (r *Repository) UpdateOneById(id int, avatar *string, isActive *bool) (*int64, error) {
	findUsr, err := r.FindOneById(id)
	if err != nil || findUsr == nil {
		return nil, err
	}

	sqlStatement := `
        UPDATE users
        SET
            avatar = COALESCE(:avatar, avatar),
            is_active = COALESCE(:is_active, is_active)
        WHERE id = (:id)
    `

	res, errUpdt := r.db.NamedExec(sqlStatement, map[string]interface{}{
		"avatar":    avatar,
		"is_active": isActive,
		"id":        id,
	})

	if errUpdt != nil {
		fmt.Println("Insert failed", errUpdt)
		return nil, errUpdt
	}

	affected, errAfct := res.RowsAffected()
	if errAfct != nil {
		fmt.Println("Retrieving affected rows failed:", errAfct)
		return nil, errAfct
	}

	return &affected, nil

}
