package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shipherman/gophermart/ent"
	"github.com/shipherman/gophermart/lib/transport/middleware"
)

// User login page
func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var u ent.User
	a := middleware.NewAuthenticator(h.Client)
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// fmt.Printf("u:%s, p:%s", u.Login, u.Password)

	jwt, err := a.Auth(u.Login, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Authorization", jwt)
	w.WriteHeader(http.StatusOK)
}
