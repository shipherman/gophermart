package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shipherman/gophermart/internal/models"
)

// Pay with bonuses
func (h *Handler) HandlePostWithdraw(w http.ResponseWriter, r *http.Request) {
	var newWithdraw models.WithdrawResponse

	// Execute user from context
	user := r.Context().Value(models.UserCtxKey{}).(string)

	err := json.NewDecoder(r.Body).Decode(&newWithdraw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Client.UpdateWithdraw(user, newWithdraw.Sum)
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

	err = h.Client.InsertWithdraw(user, newWithdraw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
