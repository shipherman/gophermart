package handlers

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/shipherman/gophermart/lib/accrual"
	"github.com/shipherman/gophermart/lib/models"
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

		h.processOrder(newOrder)
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
	defer close(errCh)

	// Register order as a new one
	newOrder.Status = models.New
	newOrder.TimeStamp = time.Now()

	go h.Client.InsertOrder(newOrder, errCh)

	go accrual.ReqAccural(newOrder.OrderNum, h.Client, errCh)

	for err := range errCh {
		log.Print("err:", err)
	}
}
