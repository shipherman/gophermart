package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shipherman/gophermart/ent"
	"github.com/shipherman/gophermart/lib/db"
	"github.com/shipherman/gophermart/lib/models"
	"github.com/shipherman/gophermart/lib/transport/middleware"

	"github.com/go-chi/chi/v5"
)

// Create handler instance
func NewHandler() {}

// Return main page
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Root page\n"))
}

// User registration page
// Generate random password; save tuple to DB. Return non-ecrypted password to user
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

	w.WriteHeader(http.StatusOK)
}

// User login page
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var u ent.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	jwt, err := middleware.Auth(u.Login, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Write([]byte(jwt))

}

// Get a list of user orders
func HandleGetOrders(w http.ResponseWriter, r *http.Request) {}

// Create a new order
func HandlePostOrders(w http.ResponseWriter, r *http.Request) {}

// Get bonuses balance
func HandleBalance(w http.ResponseWriter, r *http.Request) {
	var balance models.Balance
	var err error

	// Execute user from context
	user := chi.URLParam(r, "user")

	balance, err = db.SelectBalance(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(balance)

	w.WriteHeader(http.StatusOK)
}

// Pay with bonuses
func HandlePostWithdraw(w http.ResponseWriter, r *http.Request) {}

// Get all transactions for user bonus account
func HandleGetWithdrawals(w http.ResponseWriter, r *http.Request) {}
