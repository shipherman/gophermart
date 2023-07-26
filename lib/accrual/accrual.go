// Interaction with accural service
package accrual

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/shipherman/gophermart/lib/db"
	"github.com/shipherman/gophermart/lib/models"
	// "github.com/cenkalti/backoff/v4"
)

var addr string

func SetAccrualAddress(s string) {
	addr = s
}

// Client requests order
// Parses results allign to provided codes
// Return it to handler
func parseBody(r *resty.Response) (order *models.OrderResponse, err error) {
	err = json.Unmarshal(r.Body(), &order)
	if err != nil {
		return order, err
	}

	return order, nil
}

// Request Accural for discount
func ReqAccrual(orderResp *models.OrderResponse, dbc *db.DBClient, errCh chan error) {
	var done bool = false

	defer close(errCh)

	client := resty.New()

	// Build connection string for Accrual app
	orderAddr := fmt.Sprintf("%s/api/orders/%s", addr, orderResp.OrderNum)

	for !done {

		// Get accural for the order
		resp, err := client.R().EnableTrace().
			Get(orderAddr)
		fmt.Printf("reqAcc response: %v; Addr: %s\n", resp, orderAddr)
		if err != nil {
			errCh <- err
			return
		}

		switch resp.StatusCode() {
		// успешная обработка запроса
		case 200:
			// Parse accrual response and save to
			// OrderREsp structure
			orderParsed, err := parseBody(resp)
			orderResp.Status = orderParsed.Status
			orderResp.Accural = orderParsed.Accural

			if err != nil {
				errCh <- err
			}

			err = dbc.UpdateOrder(*orderResp)
			if err != nil {
				errCh <- err
			}

			err = dbc.UpdateBalance(*orderResp)
			if err != nil {
				errCh <- err
			}
			fmt.Println(orderResp)
			if orderResp.Status == "PROCESSED" || orderResp.Status == "INVALID" {
				done = true
			}
		// заказ не зарегистрирован в системе расчёта
		case 204:
			// orderResp.Status = "IVALID"
			// err = dbc.UpdateOrder(*orderResp)
			// if err != nil {
			// 	errCh <- err
			// }
			// done = true
		// превышено количество запросов к сервису
		case 429:
			orderResp.Status = "PROCESSING"
			err = dbc.UpdateOrder(*orderResp)
			if err != nil {
				errCh <- err
			}
		// внутренняя ошибка сервера
		case 500:
			// to do
		case 404:
			// errCh <- fmt.Errorf("accural app is not configured")
		}
		time.Sleep(5 * time.Second)
	}
}
