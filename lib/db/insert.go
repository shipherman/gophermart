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

func InsertOrder(newOrder models.Order) error {
	client := GetClient()

	// put order to accrual app
	err := acc.ReqAccural(newOrder.OrderNum)
	if err != nil {
		return err
	}

	user, err := SelectUser(newOrder.User)
	if err != nil {
		return err
	}

	//save data to db
	order, err := client.Order.Create().
		SetOrdernum(newOrder.OrderNum).
		SetStatus(newOrder.Status).
		SetAccural(33).
		SetTimestamp(time.Now()).
		SetUser(user).
		Save(context.Background())

	if err != nil {
		return err
	}

	fmt.Println(order)
	return nil
}
