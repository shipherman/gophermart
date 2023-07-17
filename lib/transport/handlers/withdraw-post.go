package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shipherman/gophermart/lib/db"
	"github.com/shipherman/gophermart/lib/models"
)

// Pay with bonuses
func HandlePostWithdraw(w http.ResponseWriter, r *http.Request) {
	var newWithdraw models.Withdraw

	// Execute user from context
	newWithdraw.User = chi.URLParam(r, "user")

	err := json.NewDecoder(r.Body).Decode(&newWithdraw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.UpdateWithdraw(newWithdraw.User, newWithdraw.Sum)
	if err != nil {
		switch err.Error() {
		case "not anough bonuses to withdraw":
			http.Error(w, err.Error(), http.StatusPaymentRequired)
			return
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	err = db.InsertWithdraw(newWithdraw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
