package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shipherman/gophermart/lib/db"
)

// Get a list of user orders
func HandleGetOrders(w http.ResponseWriter, r *http.Request) {
	u := chi.URLParam(r, "user")

	orders, err := db.SelectOrders(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "applicaion/json")
	json.NewEncoder(w).Encode(orders)
}
