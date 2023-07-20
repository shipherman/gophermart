package db

import (
	"context"
	"fmt"

	"github.com/shipherman/gophermart/ent"
	"github.com/shipherman/gophermart/ent/order"
	"github.com/shipherman/gophermart/ent/user"
	"github.com/shipherman/gophermart/ent/withdrawals"
	"github.com/shipherman/gophermart/lib/models"
)

// Get bonuses balance for provided user
func (dbc *DBClient) SelectBalance(u string) (response models.BalanceResponse, err error) {
	req, err := dbc.Client.User.
		Query().
		Where(user.LoginEQ(u)).
		All(context.Background())
	if err != nil {
		return response, err
	}

	response.Current = req[0].Balance
	response.Withdrawn = req[0].Withdraw

	return response, nil
}

// Get user by login
func (dbc *DBClient) SelectUserExistence(u, p string) (bool, error) {
	var exist = false
	req, err := dbc.Client.User.
		Query().
		Where(user.LoginEQ(u)).
		All(context.Background())
	if err != nil {
		return exist, err
	}

	if len(req) == 0 {
		return exist, fmt.Errorf("user does not exist, please register")
	}

	if req[0].Password != p {
		return exist, fmt.Errorf("wrong password")
	} else {
		exist = true
	}

	return exist, nil
}

func (dbc *DBClient) SelectUser(u string) (*ent.User, error) {
	user, err := dbc.Client.User.
		Query().
		Where(user.LoginEQ(u)).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	return user[0], nil
}

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

func (dbc *DBClient) SelectWithdrawals(u string) ([]models.WithdrawResponse, error) {
	var wsResp []models.WithdrawResponse

	entWs, err := dbc.Client.Withdrawals.
		Query().
		Where(withdrawals.HasUserWith(user.Login(u))).
		All(context.Background())
	if err != nil {
		return wsResp, err
	}

	for _, w := range entWs {
		fmt.Println()
		var wdraw models.WithdrawResponse
		wdraw.OrderNum = w.Order
		wdraw.Sum = w.Sum
		wdraw.TimeStamp = w.Timestamp
		wsResp = append(wsResp, wdraw)
	}
	return wsResp, nil
}
