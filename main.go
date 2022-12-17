package main

import (
	"golang-chi/config"
	"golang-chi/delivery/response"
	"golang-chi/delivery/router"
	"net/http"
)

func main() {
	config.Init()
	router.Routers()
	http.ListenAndServe(":8080", response.Logger())
}
