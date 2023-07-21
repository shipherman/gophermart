package db

import (
	"context"

	"github.com/shipherman/gophermart/ent/order"
	"github.com/shipherman/gophermart/ent/user"
	"github.com/shipherman/gophermart/lib/models"
)

// INSERT new order
func (dbc *DBClient) InsertOrder(newOrder models.OrderResponse, errCh chan error) {
	// // put orderResp to accrual app
	// accResp, err := accrual.ReqAccural(newOrder.OrderNum)
	// if err != nil {
	// 	errCh <- err
	// }

	// newOrder.Status = accResp.Status
	// newOrder.Accural = accResp.Accural
	// newOrder.TimeStamp = time.Now()

	// Get ent User struct
	user, err := dbc.SelectUser(newOrder.User)
	if err != nil {
		errCh <- err
	}

	// Save new Order to db
	_, err = dbc.Client.Order.Create().
		SetOrdernum(newOrder.OrderNum).
		SetStatus(newOrder.Status).
		SetAccural(newOrder.Accural).
		SetTimestamp(newOrder.TimeStamp).
		SetUser(user).
		Save(context.Background())

	if err != nil {
		errCh <- err
	}

	errCh <- err
}

// UPDATE existing order
func (dbc *DBClient) UpdateOrder(order models.OrderResponse) error {
	_, err := dbc.Client.Order.Update().
		SetStatus(order.Status).Save(context.Background())
	return err
}

// SELECT Order owner
func (dbc *DBClient) SelectOrderOwner(on int) (string, error) {
	order, err := dbc.Client.Order.
		Query().
		Where(order.OrdernumEQ(on)).
		All(context.Background())
	if err != nil || len(order) == 0 {
		return "", err
	}

	u, err := order[0].QueryUser().All(context.Background())
	if err != nil {
		return "", err
	}

	return u[0].Login, nil
}

func (dbc *DBClient) SelectOrders(u string) ([]models.OrderResponse, error) {
	var orderResp []models.OrderResponse

	entOrder, err := dbc.Client.Order.
		Query().
		Where(order.HasUserWith(user.Login(u))).
		All(context.Background())
	if err != nil {
		return orderResp, err
	}

	for _, o := range entOrder {
		var order models.OrderResponse
		order.Accural = o.Accural
		order.OrderNum = o.Ordernum
		order.Status = o.Status
		order.TimeStamp = o.Timestamp
		orderResp = append(orderResp, order)
	}

	return orderResp, nil
}
