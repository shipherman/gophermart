package db

import (
	"context"
	"fmt"

	"github.com/shipherman/gophermart/ent/order"
	"github.com/shipherman/gophermart/ent/user"
	"github.com/shipherman/gophermart/lib/models"
)

// INSERT new order
func (dbc *DBClient) InsertOrder(newOrder models.OrderResponse) error {
	// // put orderResp to accrual app
	// accResp, err := accrual.ReqAccrual(newOrder.OrderNum)
	// if err != nil {
	// 	return err
	// }

	// newOrder.Status = accResp.Status
	// newOrder.Accrual = accResp.Accrual
	// newOrder.TimeStamp = time.Now()

	// Get ent User struct
	fmt.Println(newOrder)

	user, err := dbc.SelectUser(newOrder.User)
	if err != nil {
		return err
	}

	// Save new Order to db
	_, err = dbc.Client.Order.Create().
		SetOrdernum(newOrder.OrderNum).
		SetStatus(newOrder.Status).
		SetAccrual(newOrder.Accrual).
		SetTimestamp(newOrder.TimeStamp).
		SetUser(user).
		Save(context.Background())

	if err != nil {
		return err
	}

	return nil
}

// UPDATE existing order
func (dbc *DBClient) UpdateOrder(orderResp models.OrderResponse) error {
	fmt.Println("Update order:", orderResp)
	o, err := dbc.Client.Order.
		Query().
		Where(order.OrdernumEQ(orderResp.OrderNum)).
		First(context.Background())
	if err != nil {
		return fmt.Errorf("UpdateOrder error during selecting order: %w", err)
	}

	_, err = o.Update().
		SetStatus(orderResp.Status).Save(context.Background())
	if err != nil {
		return fmt.Errorf("UpdateOrder error: %w", err)
	}

	return nil
}

// SELECT Order owner
func (dbc *DBClient) SelectOrderOwner(on string) (string, error) {
	order, err := dbc.Client.Order.
		Query().
		Where(order.OrdernumEQ(on)).
		All(context.Background())
	if err != nil || len(order) == 0 {
		return "", err
	}

	u, err := order[0].QueryUser().All(context.Background())
	if err != nil {
		return "", fmt.Errorf("SelectOrderowner error: %w", err)
	}

	return u[0].Login, nil
}

// SELECT Orders
func (dbc *DBClient) SelectOrders(u string) ([]models.OrderResponse, error) {
	var orderResp []models.OrderResponse

	entOrder, err := dbc.Client.Order.
		Query().
		Where(order.HasUserWith(user.Login(u))).
		All(context.Background())
	if err != nil {
		return orderResp, fmt.Errorf("SelectOrders error: %w", err)
	}

	for _, o := range entOrder {
		var order models.OrderResponse
		order.Accrual = o.Accrual
		order.OrderNum = o.Ordernum
		order.Status = o.Status
		order.TimeStamp = o.Timestamp
		orderResp = append(orderResp, order)
	}

	return orderResp, nil
}
