package handlers

import (
	"bytes"
	"net/http"
	"strconv"
	"time"

	"github.com/shipherman/gophermart/internal/clients"
	"github.com/shipherman/gophermart/internal/models"

	"github.com/shipherman/gophermart/pkg/luhn"
)

// Create a new order
func (h *Handler) HandlePostOrder(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	var newOrder models.OrderResponse

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save user login from context
	newOrder.User = r.Context().Value(models.UserCtxKey{}).(string)
	newOrder.OrderNum = buf.String()

	// Check if order is already registred by someone
	u, err := h.Client.SelectOrderOwner(newOrder.OrderNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// NULL user means there is no such order in DB registred
	if u == nil {
		orderInt, err := strconv.Atoi(newOrder.OrderNum)
		if !luhn.Valid(orderInt) || err != nil {
			http.Error(w, "wrong format of order number", http.StatusUnprocessableEntity)
			return
		}

		// Write order to DB
		// newOrder := models.OrderResponse{}
		newOrder.Status = models.New
		newOrder.TimeStamp = time.Now()
		err = h.Client.InsertOrder(newOrder)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		// Inform client the order has been accepted by server
		// And then get accrual amount
		w.WriteHeader(http.StatusAccepted)
		go clients.ReqAccrual(&newOrder, h.Client)

		return
	}

	// Order uploaded by current user
	// Returning Status 200
	if u.User == newOrder.User {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}
