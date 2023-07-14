package models

import (
	"time"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Order struct {
	OrderNum  int
	User      string
	Status    string
	TimeStamp time.Time
}

type Balance struct {
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}
