// Interaction with accural service
package acc

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/shipherman/gophermart/lib/models"
)

var addr string

func SetAccuralAddress(s string) {
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

func ReqAccural(orderNum int) (order models.OrderResponse, err error) {

	client := resty.New()

	// Build connection string
	addr = fmt.Sprintf("http://%s/api/order/%d", addr, orderNum)

	// Get accural for the order
	resp, err := client.R().EnableTrace().
		Get(addr)
	if err != nil {
		return order, err
	}

	switch resp.StatusCode() {
	// успешная обработка запроса
	case 200:
		order, err := parseBody(resp)
		if err != nil {
			return order, err
		}
		return order, nil
	// заказ не зарегистрирован в системе расчёта
	case 204:
		order.Status = "IVALID"
		return order, nil
	// превышено количество запросов к сервису
	case 429:
		order.Status = "PROCESSING"
		return order, nil
	// внутренняя ошибка сервера
	case 500:
	case 404:
		return order, fmt.Errorf("accural app is not configured")
	}

	fmt.Println("Accural response body: ", resp)

	return order, nil
}
