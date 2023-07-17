package db

import (
	"context"
	"fmt"
	"time"

	"github.com/shipherman/gophermart/ent"
	"github.com/shipherman/gophermart/lib/acc"
	"github.com/shipherman/gophermart/lib/models"
)

func InsertUser(newUser ent.User) error {
	client := GetClient()
	user, err := client.User.Create().
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

func InsertOrder(newOrder models.OrderResponse, errCh chan error) {

	client := GetClient()

	// // put orderResp to accrual app
	accResp, err := acc.ReqAccural(newOrder.OrderNum)
	if err != nil {
		errCh <- err
	}

	newOrder.Status = accResp.Status
	newOrder.Accural = accResp.Accural
	newOrder.TimeStamp = time.Now()

	// Get ent User struct
	user, err := SelectUser(newOrder.User)
	if err != nil {
		errCh <- err
	}

	// Save new Order to db
	_, err = client.Order.Create().
		SetOrdernum(newOrder.OrderNum).
		SetStatus(newOrder.Status).
		SetAccural(newOrder.Accural).
		SetTimestamp(newOrder.TimeStamp).
		SetUser(user).
		Save(context.Background())

	if err != nil {
		errCh <- err
	}

	errCh <- nil

}
