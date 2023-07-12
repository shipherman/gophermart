package models

import (
	"time"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Order struct {
	ID        int
	UserID    int
	Status    string
	TimeStamp time.Time
}
