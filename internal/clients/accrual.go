// Interaction with Accrual service
package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-resty/resty/v2"
	"github.com/shipherman/gophermart/internal/db"
	"github.com/shipherman/gophermart/internal/models"
)

type options struct {
	Address string
	Timeout time.Duration
}

// Accrual config
var cfg = options{}
var logEntry = middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags)}

func ConfigureAccrual(a string, t time.Duration) {
	cfg.Address = a
	cfg.Timeout = t
}

// Client requests order
// Parses results allign to provided codes
// Return it to handler
func parseBody(r *resty.Response) (order *models.OrderResponse, err error) {
	err = json.Unmarshal(r.Body(), &order)
	if err != nil {
		return order, fmt.Errorf("error during parsing to json: %w", err)
	}
	return order, nil
}

// Request Accrual for discount
func ReqAccrual(orderResp *models.OrderResponse, dbc db.DBClientInt) {
	client := resty.New()
	client.RetryMaxWaitTime = time.Second * 1
	client.RetryCount = 5

	// fmt.Println("running accrual")
	// Build connection string for Accrual app
	orderAddr := fmt.Sprintf("%s/api/orders/%s", cfg.Address, orderResp.OrderNum)

	// Create lambda to use it in backoff.Retry()
	f := func() error {
		// Get Accrual for the order
		resp, err := client.R().EnableTrace().
			Get(orderAddr)
		// fmt.Printf("resp code: %v; resp body: %v; Addr: %s\n", resp.StatusCode(), resp, orderAddr)
		if err != nil {
			return err
		}

		// fmt.Println(resp.StatusCode())

		switch resp.StatusCode() {
		// Успешная обработка запроса
		case 200:
			// Parse accrual response and save to
			// OrderREsp structure
			parsedBody, err := parseBody(resp)
			if err != nil {
				return fmt.Errorf("ReqAccrual parsing Accrual reponse error: %w", err)
			}
			orderResp.Status = parsedBody.Status
			orderResp.Accrual = parsedBody.Accrual

			err = dbc.UpdateOrder(*orderResp)
			if err != nil {
				return err
			}

			err = dbc.UpdateBalance(*orderResp)
			if err != nil {
				return err
			}
			if orderResp.Status == "PROCESSED" || orderResp.Status == "INVALID" {
				return nil
			}
		// Заказ не зарегистрирован в системе расчёта
		case 204:
			orderResp.Status = "PROCESSING"
			err = dbc.UpdateOrder(*orderResp)
			if err != nil {
				return err
			}
		// Превышено количество запросов к сервису
		case 429:
			orderResp.Status = "PROCESSING"
			err = dbc.UpdateOrder(*orderResp)
			if err != nil {
				return fmt.Errorf("too much requests error, retry in 60 sec: %w", err)
			}
		default:
			return nil
		}

		return nil
	}

	// Use backoff package to implement retryer with increasing interval between attempts
	b := backoff.NewExponentialBackOff()
	b.MaxInterval = cfg.Timeout
	err := backoff.Retry(f, b)
	if err != nil {
		logEntry.Logger.Print(fmt.Errorf("ReqAccrual error: %w", err))
	}

	log.Println(orderResp.OrderNum, " order is processed by accrual app")

}
