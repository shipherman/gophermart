package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shipherman/gophermart/internal/models"
)

// Get a list of user orders
func (h *Handler) HandleGetOrders(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(models.UserCtxKey{}).(string)

	orders, err := h.Client.SelectOrders(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}
