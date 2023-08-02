package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shipherman/gophermart/internal/models"
	"github.com/shipherman/gophermart/internal/transport/middleware"
)

// Get bonuses balance
func (h *Handler) HandleBalance(w http.ResponseWriter, r *http.Request) {
	var balance models.BalanceResponse
	var err error

	// Execute user from context
	user := r.Context().Value(middleware.UserCtxKey{}).(string)

	log.Println(user)
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
