package handlers

import (
	"github.com/EggsyOnCode/go-rest-api-server/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

//H capital in the internal package indicates that it can be called from outside(public func) instead of a private one

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	//registering a route such as express.app().route('///')
	r.Route("/accounts", func(router chi.Router) {

		router.Use(middleware.Authorization)
		// this is equivalent ot registering a GET endpoint in the endpoint '/accounts'
		router.Get("/coins", GetCoinBalance)
	})
}
