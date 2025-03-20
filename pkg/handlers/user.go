package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jaibhavaya/dashboard-go/pkg/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func GetUsersHandler(userRepo *models.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := userRepo.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"status": "success",
			"data":   users,
		})
	}
}

func GetUserByIDHandler(userRepo *models.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, err := userRepo.FindByID(id)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"status": "success",
			"data":   user,
		})
	}
}

func GetAddHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		xStr := r.URL.Query().Get("x")
		yStr := r.URL.Query().Get("y")

		x, err := strconv.Atoi(xStr)
		if err != nil {
			http.Error(w, "Invalid value for x", http.StatusBadRequest)
			return
		}

		y, err := strconv.Atoi(yStr)
		if err != nil {
			http.Error(w, "Invalid value for y", http.StatusBadRequest)
			return
		}

		result := x + y

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"status": "success",
			"data":   result,
		})
	}
}
