package db

import (
	"context"
	"fmt"
	"time"

	"github.com/shipherman/gophermart/ent"
	"github.com/shipherman/gophermart/lib/accrual"
	"github.com/shipherman/gophermart/lib/models"
)

func (dbc *DBClient) InsertUser(newUser ent.User) error {
	user, err := dbc.Client.User.Create().
		SetLogin(newUser.Login).
		SetPassword(newUser.Password).
		SetBalance(0).
		SetWithdraw(0).
		Save(context.Background())

	fmt.Println(user, err)
	if err != nil {
		return err
	}
	return nil
}

func (dbc *DBClient) InsertOrder(newOrder models.OrderResponse) error {
	// // put orderResp to accrual app
	accResp, err := accrual.ReqAccural(newOrder.OrderNum)
	if err != nil {
		return err
	}

	newOrder.Status = accResp.Status
	newOrder.Accural = accResp.Accural
	newOrder.TimeStamp = time.Now()

	// Get ent User struct
	user, err := dbc.SelectUser(newOrder.User)
	if err != nil {
		return err
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
		return err
	}

	return nil

}

func (dbc *DBClient) InsertWithdraw(u string, newWithdraw models.WithdrawResponse) error {
	user, err := dbc.SelectUser(u)
	if err != nil {
		return err
	}

	newWithdraw.TimeStamp = time.Now()
	fmt.Println(newWithdraw.TimeStamp.String())

	_, err = dbc.Client.Withdrawals.Create().
		SetOrder(newWithdraw.OrderNum).
		SetSum(newWithdraw.Sum).
		SetTimestamp(newWithdraw.TimeStamp).
		SetUser(user).
		Save(context.Background())

	if err != nil {
		return err
	}

	return nil
}
