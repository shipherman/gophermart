package db

import (
	"context"
	"fmt"

	"github.com/shipherman/gophermart/ent/user"
)

func UpdateWithdraw(u string, a int) error {
	client := GetClient()

	uent, err := client.User.Query().
		Where(user.Login(u)).First(context.Background())
	if err != nil {
		return err
	}

	if uent.Balance < a {
		return fmt.Errorf("not anough bonuses to withdraw")
	}
	uent, err = uent.Update().
		SetBalance(uent.Balance - a).
		SetWithdraw(uent.Withdraw + a).Save(context.Background())

	if err != nil {
		return err
	}

	return nil
}
