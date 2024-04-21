package handlers

import (
	"github.com/4shb0rne/goapi-basic/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes) //ignore extra slash
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
	r.Route("/product", func(router chi.Router) {
		router.Get("/view", GetAllProducts)
		router.Post("/insert", InsertProduct)
	})
}
