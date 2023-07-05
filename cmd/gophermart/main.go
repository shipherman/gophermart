package main

import (
	"net/http"

	"github.com/shipherman/gophermart/lib/transport/routes"
	//"github.com/go-chi/chi/v5"
)

func main() {
	router := routes.NewRouter()
	http.ListenAndServe(":9090", router)

}
