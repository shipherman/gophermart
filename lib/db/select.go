package db

import (
	"context"
	"fmt"

	"github.com/shipherman/gophermart/ent"
	"github.com/shipherman/gophermart/ent/user"
	"github.com/shipherman/gophermart/lib/models"
)

// Get bonuses balance for provided user
func SelectBalance(u string) (response models.BalanceResponse, err error) {
	client := GetClient()
	req, err := client.User.
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
func SelectUserExistence(u, p string) (bool, error) {
	var exist = false
	client := GetClient()
	req, err := client.User.
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

func SelectUser(u string) (*ent.User, error) {
	client := GetClient()
	user, err := client.User.
		Query().
		Where(user.LoginEQ(u)).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	if len(user) == 0 {
		return nil, fmt.Errorf("user does not exitst, please register")
	}

	return user[0], nil
}
