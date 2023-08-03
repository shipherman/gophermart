package worker

import (
	"fmt"
	"log"
	"time"

	"github.com/shipherman/gophermart/internal/clients"
	"github.com/shipherman/gophermart/internal/db"
)

type Worker struct {
	InCh    chan string
	CloseCh chan bool
	ErrCh   chan error
	Client  db.DBClientInt
}

func New(dbc db.DBClientInt) *Worker {
	in := make(chan string, 5)
	close := make(chan bool, 1)
	err := make(chan error)
	return &Worker{InCh: in, CloseCh: close, ErrCh: err, Client: dbc}
}

func (w *Worker) Run() {
	defer close(w.InCh)

	for {
		fmt.Println(time.Now())
		orderResp, err := w.Client.SelectFirstUnprocessedOrder()
		if err != nil {
			// if err == db.ErrorOrderNotFound {
			// 	continue
			// }
			w.ErrCh <- fmt.Errorf("Worker error: %w", err)
		}

		clients.ReqAccrual(&orderResp, w.Client, w.ErrCh)
		if <-w.CloseCh {
			log.Println("Closing worker goroutine")
			return
		}
	}
}
