// Interaction with accural service
package acc

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Client requests order
// Parses results allign to provided codes
// Return it to handler
func ReqAccural(orderNum int) error {
	client := resty.New()

	// Build connection string
	addr := fmt.Sprintf("http://localhost:8080/api/order/%d", orderNum)

	// Get accural for the order
	resp, err := client.R().EnableTrace().
		Get(addr)
	if err != nil {
		return err
	}

	fmt.Println("Accural response body: ", resp)

	return nil
}
