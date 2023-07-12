package db

import (
	"context"
	"fmt"

	"github.com/shipherman/gophermart/ent"
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

func InsertOrder(newOrder ent.Order) error {
	// client := GetClient()
	// put order to accrual app

	// save data to db
	// order, err := client.Order.Create().
	// 	SetOrdernum(newOrder.Ordernum).
	// 	SetStatus(newOrder.Status).
	// 	SetUser()

	return nil
}
