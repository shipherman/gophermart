package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shipherman/gophermart/ent"
	"github.com/shipherman/gophermart/lib/transport/middleware"
)

// User login page
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var u ent.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	jwt, err := middleware.Auth(u.Login, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte(jwt))

}
