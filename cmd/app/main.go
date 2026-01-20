package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	fmt.Println("Server is running on port: 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
