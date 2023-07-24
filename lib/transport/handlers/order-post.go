package handlers

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/shipherman/gophermart/lib/accrual"
	"github.com/shipherman/gophermart/lib/luhn"
	"github.com/shipherman/gophermart/lib/models"

	"github.com/go-chi/chi/v5/middleware"
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

	newOrder.User = chi.URLParam(r, "user")
	newOrder.OrderNum = buf.String()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.Client.SelectOrderOwner(newOrder.OrderNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch u {
	case "":
		orderInt, err := strconv.Atoi(newOrder.OrderNum)
		if !luhn.Valid(orderInt) || err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusAccepted)

		go h.processOrder(newOrder, r)
		return
	case newOrder.User:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusConflict)
		return
	}
}

func (h *Handler) processOrder(newOrder models.OrderResponse, r *http.Request) {
	errCh := make(chan error)

	// Logger for outgoing requests
	logEntry := middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: false}

	// Register order as a new one
	newOrder.Status = models.New
	newOrder.TimeStamp = time.Now()

	err := h.Client.InsertOrder(newOrder)
	if err != nil {
		logEntry.Logger.Print(err)
	}

	go accrual.ReqAccural(newOrder, h.Client, errCh)

	for err := range errCh {
		if err != nil {
			logEntry.Logger.Print(err)
		}
	}
}
