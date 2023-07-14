package models

import (
	"time"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type OrderResponse struct {
	OrderNum  int       `json:"number"`
	User      string    `json:"-"`
	Status    string    `json:"status"`
	Accural   int       `json:"accural"`
	TimeStamp time.Time `json:"uploaded_at"`
}

type BalanceResponse struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}
