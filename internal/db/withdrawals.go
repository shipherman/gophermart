package db

import (
	"context"
	"fmt"
	"time"

	"github.com/shipherman/gophermart/generated/ent/user"
	"github.com/shipherman/gophermart/generated/ent/withdrawals"
	"github.com/shipherman/gophermart/internal/models"
)

func (dbc *DBClient) InsertWithdraw(u string, newWithdraw models.WithdrawResponse) error {
	user, err := dbc.SelectUser(u)
	if err != nil {
		return err
	}

	newWithdraw.TimeStamp = time.Now()

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

func (dbc *DBClient) UpdateWithdraw(u string, a float64) error {
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
		var wdraw models.WithdrawResponse
		wdraw.OrderNum = w.Order
		wdraw.Sum = w.Sum
		wdraw.TimeStamp = w.Timestamp
		wsResp = append(wsResp, wdraw)
	}
	return wsResp, nil
}
