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

func NewTodoRoute(db *sql.DB, router *mux.Router) {
	repo := repository.NewTodoRepository(db)
	s := services.NewTodoService(repo)
	c := controllers.NewTodoController(s)

	router.Use(middlewares.AuthMiddleware)
	router.HandleFunc("/", api.MakeHandlerFunc(c.GetAll)).Methods("GET")
	router.HandleFunc("/{id}", api.MakeHandlerFunc(c.GetById)).Methods("GET")
	router.HandleFunc("/", api.MakeHandlerFunc(c.Create)).Methods("POST")
	router.HandleFunc("/{id}", api.MakeHandlerFunc(c.Delete)).Methods("DELETE")
	router.HandleFunc("/{id}", api.MakeHandlerFunc(c.UpdateOne)).Methods("PATCH")
}
