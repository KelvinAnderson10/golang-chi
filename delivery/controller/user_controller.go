package controller

import (
	"encoding/json"
	"golang-chi/config"
	"golang-chi/delivery/response"
	"golang-chi/model"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	json.NewDecoder(r.Body).Decode(&user)

	result := config.DB.Create(&user)
	if result.Error != nil {
		response.RespondWithError(w, result.Error)
		return
	}
	response.RespondwithJSON(w, user)
}

func GetUserDetail(w http.ResponseWriter, r *http.Request) {
	var user *model.User
	id := chi.URLParam(r, "id")
	result := config.DB.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		response.RespondWithError(w, result.Error)
		return
	}
	response.RespondwithJSON(w, user)
}
