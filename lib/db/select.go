package db

import (
	"context"

	"github.com/shipherman/gophermart/ent/user"
	"github.com/shipherman/gophermart/lib/models"
)

// Get bonuses balance for provided user
func SelectBalance(u string) (response models.Balance, err error) {
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

	if req[0].Password != p {
		return exist, nil
	} else {
		exist = true
	}

	return exist, nil
}
