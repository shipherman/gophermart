package db

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
)

func InsertWithDraw() {}

func InsertUser() {
	suf := rand.New(rand.NewSource(9999))
	client := GetClient()
	user, err := client.User.Create().
		SetLogin("login" + strconv.Itoa(suf.Int())).
		SetPassword("pass").
		SetBalance(10).
		SetWithdraw(0).
		Save(context.Background())

	fmt.Println(user, err)
}
