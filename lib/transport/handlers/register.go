package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shipherman/gophermart/generated/ent"
	"github.com/shipherman/gophermart/lib/transport/middleware"
)

// User registration page
func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var newUser ent.User

	body := r.Body
	err := json.NewDecoder(body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Client.InsertUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Authenticator instance
	a := middleware.NewAuthenticator(h.Client)

	// Generate JWT
	jwt, err := a.Auth(newUser.Login, newUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Authorization", jwt)
	w.WriteHeader(http.StatusOK)
}
