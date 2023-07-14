package models

import (
	"time"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Order struct {
	User     string
	OrderNum int
	Status   string
}
type OrderResponse struct {
	OrderNum  int       `json:"number"`
	Status    string    `json:"status"`
	Accural   int       `json:"accural"`
	TimeStamp time.Time `json:"uploaded_at"`
}

type BalanceResponse struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}
