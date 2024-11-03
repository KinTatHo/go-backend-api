package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/kintatho/go-backend-api/internal/middleware"
)
// before primary function
func Handler(r *chi.Mux) {
	// global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		GetCoinBalance := 
		router.Get("/coins", GetCoinBalance)
	})
}

// GetCoinBalance is a handler for getting the coin balance of a user
