package routes

import (
	"github.com/shipherman/gophermart/lib/transport/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", handlers.HandleRoot)
	router.Post("/api/user/register", handlers.HandleRegister)
	router.Post("/api/user/login", handlers.HandleLogin)
	router.Post("/api/user/orders", handlers.HandlePostOrders)
	router.Get("/api/user/orders", handlers.HandleGetOrders)
	router.Get("/api/user/balance", handlers.HandleBalance)
	router.Post("/api/user/balance/withdraw", handlers.HandlePostWithdraw)
	router.Get("/api/user/withdrawals", handlers.HandleGetWithdrawals)

	return router
}
