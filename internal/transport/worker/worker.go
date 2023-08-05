package worker

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/shipherman/gophermart/internal/clients"
	"github.com/shipherman/gophermart/internal/db"
)

type Worker struct {
	CloseCh chan bool
	Client  db.DBClientInt
}

var logEntry = middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags)}

func New(dbc db.DBClientInt) *Worker {
	close := make(chan bool, 1)
	return &Worker{CloseCh: close, Client: dbc}
}

func (w *Worker) Run(wg *sync.WaitGroup) {
	wg.Add(1)
	for {
		time.Sleep(time.Second * 10)
		// fmt.Println("run worker")
		orderResp, err := w.Client.SelectFirstUnprocessedOrder()
		if err != nil {
			logEntry.Logger.Print(fmt.Errorf("Worker error: %w", err))
		} else {
			clients.ReqAccrual(&orderResp, w.Client)
		}

		select {
		case <-w.CloseCh:
			log.Println("Closing worker goroutine")
			wg.Done()
			return
		default:
			continue
		}
	}
}
