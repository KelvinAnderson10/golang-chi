package router

import (
	"golang-chi/config"
	"golang-chi/delivery/controller"

	"github.com/go-chi/chi/v5"
)

func Routers() *chi.Mux {
	router := config.Router
	router.Get("/user/{id}", controller.GetUserDetail)
	router.Post("/user", controller.CreateUser)

	return router
}
