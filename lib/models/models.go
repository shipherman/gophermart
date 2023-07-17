package models

import (
	"time"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type WithdrawResponse struct {
	OrderNum  int       `json:"order"`
	Sum       int       `json:"sum"`
	TimeStamp time.Time `json:"processed_at"`
}

type OrderResponse struct {
	OrderNum  int       `json:"number"`
	User      string    `json:"-"`
	Status    string    `json:"status"`
	Accural   int       `json:"accural,omitempty"`
	TimeStamp time.Time `json:"uploaded_at"`
}

type BalanceResponse struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}
