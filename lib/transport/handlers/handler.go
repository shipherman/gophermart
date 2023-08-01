package handlers

import "github.com/shipherman/gophermart/lib/db"

type Handler struct {
	Client db.DBClientInt
}

// Create handler instance ???
func NewHandler(dbclient db.DBClientInt) *Handler {
	return &Handler{Client: dbclient}
}
