package models

import (
	"time"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type WithdrawResponse struct {
	OrderNum  string    `json:"order"`
	Sum       int       `json:"sum"`
	TimeStamp time.Time `json:"processed_at,omitempty"`
}

type OrderResponse struct {
	OrderNum  string    `json:"number"`
	User      string    `json:"-"`
	Status    string    `json:"status"`
	Accural   int       `json:"accural,omitempty"`
	TimeStamp time.Time `json:"uploaded_at"`
}

type BalanceResponse struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}

const (
	New        = "NEW"
	Processing = "PROCESSING"
	Invalid    = "INVALID"
	Done       = "PROCESSED"
)
