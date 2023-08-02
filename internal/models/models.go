package models

import (
	"time"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Structure to provide username throuh context to handlers
type UserCtxKey struct{}

type WithdrawResponse struct {
	OrderNum  string    `json:"order"`
	Sum       float64   `json:"sum"`
	TimeStamp time.Time `json:"processed_at,omitempty"`
}

type OrderResponse struct {
	OrderNum  string    `json:"number"`
	User      string    `json:"-"`
	Status    string    `json:"status"`
	Accrual   float64   `json:"accrual,omitempty"`
	TimeStamp time.Time `json:"uploaded_at,omitempty"`
}

type BalanceResponse struct {
	Current   float64 `json:"current"`
	Withdrawn float64 `json:"withdrawn"`
}

const (
	New        = "NEW"
	Processing = "PROCESSING"
	Invalid    = "INVALID"
	Done       = "PROCESSED"
)
