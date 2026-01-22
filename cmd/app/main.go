package main

import (
	"encoding/json"
	"finance-app-be/internal/db"
	"finance-app-be/internal/users/repository"
	"finance-app-be/internal/users/service"
	"finance-app-be/schema"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func main() {
	connect := db.Connection{}
	db, err := connect.Connect()
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}

	// migrate(db)
	defer db.Close()

	r := chi.NewRouter()
	routes(r, db)
	r.Mount("/api", r)

	fmt.Println("Server is running on port: 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func routes(r *chi.Mux, db *sqlx.DB) {
	// USERS MODULE
	userRepo := repository.NewRepository(db)
	userService := service.NewService(userRepo)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			users, err := userService.FindAll()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
		})

		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.Atoi(chi.URLParam(r, "id"))
			if id == 0 || err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}

			user, errFind := userService.FindOneById(id)
			if errFind != nil {
				http.Error(w, errFind.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var reqBody struct {
				Email    string `json:"email"`
				Password string `json:"password"`
			}

			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if reqBody.Email == "" || reqBody.Password == "" {
				http.Error(w, "Email or Password is empty", http.StatusBadRequest)
				return
			}

			rowsAffected, err := userService.Create(reqBody.Email, reqBody.Password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct {
				RowsAffected *int64 `json:"rows_affected"`
			}{RowsAffected: rowsAffected})
		})

		r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
			var reqBody struct {
				Avatar   *string `json:"avatar"`
				IsActive *bool   `json:"is_active"`
			}

			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			id, err := strconv.Atoi(chi.URLParam(r, "id"))
			if id == 0 || err != nil {
				http.Error(w, "Invalid user ID", http.StatusBadRequest)
				return
			}

			rowsAffected, err := userService.UpdateOneById(id, reqBody.Avatar, reqBody.IsActive)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(struct {
				RowsAffected *int64 `json:"rows_affected"`
			}{RowsAffected: rowsAffected})

		})

	})
}

func migrate(db *sqlx.DB) {
	usersSchema := &schema.UsersSchema{}
	organizationsSchema := &schema.OrganizationsSchema{}
	branchesSchema := &schema.BranchesSchema{}
	initialSchema := &schema.InitialBalancesSchema{}
	transactionsSchema := &schema.TransactionsSchema{}

	usersSchema.Up(db)
	organizationsSchema.Up(db)
	branchesSchema.Up(db)
	initialSchema.Up(db)
	transactionsSchema.Up(db)
}
