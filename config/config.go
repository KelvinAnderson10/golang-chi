package config

import (
	"fmt"
	"golang-chi/model"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Router *chi.Mux

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = "Enigma123"
	DB_NAME     = "chi_crud"
)

func Init() {
	Router = chi.NewRouter()
	Router.Use(middleware.Recoverer)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db.Debug()
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
