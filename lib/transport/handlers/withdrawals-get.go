package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shipherman/gophermart/lib/db"
	"github.com/shipherman/gophermart/lib/models"
)

// Get all transactions for user bonus account
func HandleGetWithdrawals(w http.ResponseWriter, r *http.Request) {
	var wsResp []models.WithdrawResponse

	user := chi.URLParam(r, "user")

	wsResp, err := db.SelectWithdrawals(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(wsResp) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = json.NewEncoder(w).Encode(wsResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
