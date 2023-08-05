package handlers

import "net/http"

// Return main page
func (h *Handler) HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Root page\n"))
}
