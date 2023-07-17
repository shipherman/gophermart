package models

import (
	"time"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Withdraw struct {
	User     string `json:"user"`
	OrderNum int    `json:"order"`
	Sum      int    `json:"sum"`
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
