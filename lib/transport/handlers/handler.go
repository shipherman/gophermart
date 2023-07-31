package handlers

import "github.com/shipherman/gophermart/lib/db"

type Handler struct {
	Client *db.DBClient
}

// Create handler instance ???
func NewHandler(dbclient *db.DBClient) *Handler {
	return &Handler{Client: dbclient}
}
