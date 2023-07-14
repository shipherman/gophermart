// Interaction with accural service
package acc

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client requests order
// Parses results
// Return it to handler
func ReqAccural() error {
	client := resty.New()

	resp, err := client.R().EnableTrace().
		Get("http://localhost:8080/api/order/134")
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}
