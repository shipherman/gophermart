package routes

import (
	"github.com/shipherman/gophermart/lib/transport/handlers"
	mid "github.com/shipherman/gophermart/lib/transport/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(h *handlers.Handler, a *mid.Authenticator) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Compress(1, "text/*", "application/*"))
	router.Get("/", h.HandleRoot)
	router.Post("/api/user/register", h.HandleRegister)
	router.Post("/api/user/login", h.HandleLogin)
	router.Post("/api/user/orders", a.CheckAuth(h.HandlePostOrder))
	router.Get("/api/user/orders", a.CheckAuth(h.HandleGetOrders))
	router.Get("/api/user/balance", a.CheckAuth(h.HandleBalance))
	router.Post("/api/user/balance/withdraw", a.CheckAuth(h.HandlePostWithdraw))
	router.Get("/api/user/balance/withdrawals", a.CheckAuth(h.HandleGetWithdrawals))

	return router
}
