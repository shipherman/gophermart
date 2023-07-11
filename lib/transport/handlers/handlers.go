package handlers

import (
	"net/http"

	"github.com/shipherman/gophermart/lib/db"
)

// Create handler instance
func NewHandler() {}

// Return main page
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Root page\n"))
	w.WriteHeader(http.StatusOK)
}

// User registration page
// Generate random password; save tuple to DB. Return non-ecrypted password to user
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	db.InsertUser()
	w.WriteHeader(http.StatusOK)
}

// User login page
func HandleLogin(w http.ResponseWriter, r *http.Request) {} // Use basic auth from chi?

// Get a list of user orders
func HandleGetOrders(w http.ResponseWriter, r *http.Request) {}

// Create a new order
func HandlePostOrders(w http.ResponseWriter, r *http.Request) {}

// Get bonuses balance
func HandleBalance(w http.ResponseWriter, r *http.Request) {}

// Pay with bonuses
func HandlePostWithdraw(w http.ResponseWriter, r *http.Request) {}

// Get all transactions for user bonus account
func HandleGetWithdrawals(w http.ResponseWriter, r *http.Request) {}
