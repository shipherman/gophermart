package handlers

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/shipherman/gophermart/lib/db"
	"github.com/shipherman/gophermart/lib/models"
)

// Create a new order
func HandlePostOrder(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var newOrder models.OrderResponse

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newOrder.User = chi.URLParam(r, "user")
	newOrder.OrderNum, err = strconv.Atoi(buf.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := db.SelectOrderOwner(newOrder.OrderNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch u {
	case "":
		w.WriteHeader(http.StatusAccepted)
		go db.InsertOrder(newOrder)
	case newOrder.User:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusConflict)
		return
	}

}
