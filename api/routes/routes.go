package routes

import (
	. "github.com/esvas/FinalProject/api/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func CreateRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/", func(r chi.Router) {
		r.Get("/", ConnectionHandler)
	})

	return r
}