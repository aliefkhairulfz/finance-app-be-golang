package main

import (
	"database/sql"
	"encoding/json"
	"finance-app-be/internal/db"
	"finance-app-be/internal/users/repository"
	"finance-app-be/internal/users/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// DB CONNECTION
	connect := db.Connection{}
	db, err := connect.Connect()
	if err != nil {
		fmt.Println("Database connection failed:", err)
		return
	}

	defer db.Close()

	r := chi.NewRouter()
	AppRoutes(r, db)

	// API PREFIX
	r.Mount("/api", r)

	// START SERVER
	fmt.Println("Server is running on port: 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func AppRoutes(r *chi.Mux, db *sql.DB) {
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
	})
}
