// Interaction with accural service
package accrual

import (
	"encoding/json"
	"fmt"

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
func parseBody(r *resty.Response) (order models.OrderResponse, err error) {
	err = json.Unmarshal(r.Body(), order)
	if err != nil {
		return order, err
	}

	return order, nil
}

func ReqAccural(orderResp models.OrderResponse, dbc *db.DBClient, errCh chan error) {
	var order models.OrderResponse

	defer close(errCh)

	client := resty.New()

	// Build connection string for Accrual app
	addr = fmt.Sprintf("http://%s/api/order/%s", addr, orderResp.OrderNum)

	// Get accural for the order
	resp, err := client.R().EnableTrace().
		Get(addr)
	// fmt.Println("reqAcc:", err)
	if err != nil {
		errCh <- err
		return
	}

	switch resp.StatusCode() {
	// успешная обработка запроса
	case 200:
		order, err := parseBody(resp)
		if err != nil {
			errCh <- err
		}

		err = dbc.UpdateOrder(order)
		if err != nil {
			errCh <- err
		}

		orderResp.Accural = order.Accural
		err = dbc.UpdateBalance(orderResp)
		if err != nil {
			errCh <- err
		}
	// заказ не зарегистрирован в системе расчёта
	case 204:
		order.Status = "IVALID"
		err = dbc.UpdateOrder(order)
		if err != nil {
			errCh <- err
		}
	// превышено количество запросов к сервису
	case 429:
		order.Status = "PROCESSING"
		err = dbc.UpdateOrder(order)
		if err != nil {
			errCh <- err
		}
	// внутренняя ошибка сервера
	case 500:
	case 404:
		errCh <- fmt.Errorf("accural app is not configured")
	}
}
