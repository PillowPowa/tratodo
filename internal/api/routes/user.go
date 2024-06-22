package routes

import (
	"database/sql"
	"tratodo/internal/api/controllers"
	"tratodo/internal/api/middlewares"
	"tratodo/internal/repository"
	"tratodo/internal/services"
	"tratodo/pkg/api"

	"github.com/gorilla/mux"
)

func NewUserRoute(db *sql.DB, router *mux.Router) {
	repo := repository.NewUserRepository(db)
	s := services.NewUserService(repo)
	c := controllers.NewUserController(s)

	pRouter := router.PathPrefix("/").Subrouter()
	pRouter.Use(middlewares.AuthMiddleware)
	pRouter.HandleFunc("/", api.MakeHandlerFunc(c.GetMe)).Methods("GET")
}
