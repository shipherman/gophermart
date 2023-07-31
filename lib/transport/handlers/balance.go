package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shipherman/gophermart/lib/models"
)

// Get bonuses balance
func (h *Handler) HandleBalance(w http.ResponseWriter, r *http.Request) {
	var balance models.BalanceResponse
	var err error

	// Execute user from context
	user := chi.URLParam(r, "user")

	balance, err = h.Client.SelectBalance(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(balance)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
