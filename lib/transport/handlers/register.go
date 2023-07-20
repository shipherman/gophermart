package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shipherman/gophermart/ent"
)

// User registration page
func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var newUser ent.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Client.InsertUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.RequestURI)

	http.Redirect(w, r, "/api/user/login", http.StatusTemporaryRedirect)
}
