package models

import (
	"time"
)

type User struct {
	ID       int
	Login    string
	Password string
	Balance  int
	Withdraw int
}

type Order struct {
	ID        int
	UserID    int
	Status    string
	TimeStamp time.Time
}
