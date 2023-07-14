package routes

import (
	"github.com/shipherman/gophermart/lib/transport/handlers"
	mid "github.com/shipherman/gophermart/lib/transport/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Compress(1, "text/*", "application/*"))
	router.Get("/", handlers.HandleRoot)
	router.Post("/api/user/register", handlers.HandleRegister)
	router.Post("/api/user/login", handlers.HandleLogin)
	router.Post("/api/user/orders", mid.CheckAuth(handlers.HandlePostOrders))
	router.Get("/api/user/orders", mid.CheckAuth(handlers.HandleGetOrders))
	router.Get("/api/user/balance", mid.CheckAuth(handlers.HandleBalance))
	router.Post("/api/user/balance/withdraw", mid.CheckAuth(handlers.HandlePostWithdraw))
	router.Get("/api/user/withdrawals", mid.CheckAuth(handlers.HandleGetWithdrawals))

	return router
}
