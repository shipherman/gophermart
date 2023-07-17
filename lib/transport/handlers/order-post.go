package handlers

import (
	"bytes"
	"log"
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
		// как отделить Хэндлер от запроса в базу и в Acural?
		go writeOrderToDB(newOrder)
	case newOrder.User:
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusConflict)
		return
	}

}

func writeOrderToDB(newOrder models.OrderResponse) {
	errCh := make(chan error)

	for err := range errCh {
		if err != nil {
			log.Println(err)
		}
	}
}
