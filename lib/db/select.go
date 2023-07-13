package db

import (
	"context"

	"github.com/shipherman/gophermart/ent/user"
)

// Get bonuses balance for provided user
func SelectBalance(user string) {}

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
