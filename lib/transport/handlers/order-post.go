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
	newOrder.OrderNum, err = strconv.Atoi(buf.String())
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
		w.WriteHeader(http.StatusAccepted)

		// как отделить Хэндлер от запроса в базу и в Acural?
		// maxlyaptsev Jul 18, 2023

		// а зачем? Получили новый заказ - сохранили, баллы можно посчитать и позже

		go h.processOrder(newOrder)
		return
	case newOrder.User:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusConflict)
		return
	}
}

func (h *Handler) processOrder(newOrder models.OrderResponse) {
	errCh := make(chan error)

	logEntry := middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: false}

	// Register order as a new one
	newOrder.Status = models.New
	newOrder.TimeStamp = time.Now()

	err := h.Client.InsertOrder(newOrder)
	if err != nil {
		logEntry.Logger.Print(err)
	}

	go accrual.ReqAccural(newOrder.OrderNum, h.Client, errCh)

	for err := range errCh {
		if err != nil {
			logEntry.Logger.Print(err)
		}
	}
}
