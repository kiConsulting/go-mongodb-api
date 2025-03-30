package handlers

import (
	"encoding/json"
	"go-mongodb-api/internal/repository"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(client *mongo.Client) *UserHandler {
	return &UserHandler{
		repo: repository.NewUserRepository(client, "centivo", "users"),
	}
}


func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.GetUsers()
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "qpplicayion/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUerIDAndAge(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	age := r.PathValue("age")
	num, err := strconv.Atoi(age)
	if err != nil {
		num = 21
	}

	// Retrieve the user from the repository
	user, err := h.repo.GetUserByIDAndAge(id, num)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found or doesn't meet age criteria", http.StatusNotFound)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
