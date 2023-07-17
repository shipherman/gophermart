package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shipherman/gophermart/ent"
	"github.com/shipherman/gophermart/lib/db"
)

// User registration page
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var newUser ent.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.InsertUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println(r.RequestURI)

	http.Redirect(w, r, "/api/user/login", http.StatusTemporaryRedirect)
}
