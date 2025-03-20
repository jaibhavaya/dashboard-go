package main

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/jaibhavaya/dashboard-go/pkg/handlers"
	"github.com/jaibhavaya/dashboard-go/pkg/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func setupRoutes(userRepo *models.UserRepository) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Get("/users", handlers.GetUsersHandler(userRepo))
		r.Get("/add", handlers.GetAddHandler())
		r.Get("/users/{id}", handlers.GetUserByIDHandler(userRepo))

		r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
			var user models.User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := userRepo.Create(user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]any{
				"status": "success",
				"data":   user,
			})
		})
	})

	// Static file server for the React build
	staticDir := filepath.Join(".", "build")
	fileServer := http.FileServer(http.Dir(staticDir))
	r.Handle("/*", http.StripPrefix("/", fileServer))

	return r
}
