package db

import (
	"context"
	"fmt"

	"github.com/shipherman/gophermart/ent/user"
)

func (dbc *DBClient) UpdateWithdraw(u string, a int) error {
	uent, err := dbc.Client.User.Query().
		Where(user.Login(u)).First(context.Background())
	if err != nil {
		return err
	}

	if uent.Balance < a {
		return fmt.Errorf("not anough bonuses to withdraw")
	}
	_, err = uent.Update().
		SetBalance(uent.Balance - a).
		SetWithdraw(uent.Withdraw + a).Save(context.Background())

	if err != nil {
		return err
	}

	return nil
}
