package db

import (
	"context"
	"fmt"

	"github.com/shipherman/gophermart/ent"
	"github.com/shipherman/gophermart/ent/user"
	"github.com/shipherman/gophermart/lib/models"
)

func (dbc *DBClient) InsertUser(newUser ent.User) error {
	_, err := dbc.Client.User.Create().
		SetLogin(newUser.Login).
		SetPassword(newUser.Password).
		SetBalance(0).
		SetWithdraw(0).
		Save(context.Background())

	return err
}

// Get user by login
func (dbc *DBClient) SelectUserExistence(u, p string) (bool, error) {
	var exist = false
	fmt.Println(u, p)
	user, err := dbc.Client.User.
		Query().
		Where(user.LoginEQ(u)).
		First(context.Background())
	if err != nil {
		return exist, err
	}

	if user == nil {
		return exist, fmt.Errorf("user does not exist, please register")
	}

	if user.Password != p {
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
