package worker

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/shipherman/gophermart/internal/clients"
	"github.com/shipherman/gophermart/internal/db"
)

type Worker struct {
	CloseCh chan bool
	ErrCh   chan error
	Client  db.DBClientInt
}

func New(dbc db.DBClientInt) *Worker {
	close := make(chan bool, 1)
	err := make(chan error)
	return &Worker{CloseCh: close, ErrCh: err, Client: dbc}
}

func (w *Worker) Run(wg *sync.WaitGroup) {
	for {
		time.Sleep(time.Second * 10)
		// fmt.Println("run worker")
		orderResp, err := w.Client.SelectFirstUnprocessedOrder()
		if err != nil {
			w.ErrCh <- fmt.Errorf("Worker error: %w", err)
		} else {
			clients.ReqAccrual(&orderResp, w.Client, w.ErrCh)
		}

		select {
		case <-w.CloseCh:
			log.Println("Closing worker goroutine")
			close(w.ErrCh)
			wg.Done()
			return
		default:
			continue
		}
	}
}
