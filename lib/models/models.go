package models

import (
	"time"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Order struct {
	OrderNum  int       `json:"number"`
	Status    string    `json:"status"`
	Accural   int       `json:"accural"`
	TimeStamp time.Time `json:"uploaded_at"`
}

type Balance struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}
