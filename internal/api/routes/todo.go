package routes

import (
	"database/sql"
	"tratodo/internal/api/controllers"
	"tratodo/internal/repository"
	"tratodo/internal/services"
	"tratodo/pkg/api"

	"github.com/gorilla/mux"
)

func NewTodoRoute(db *sql.DB, router *mux.Router) {
	repo := repository.NewTodoRepository(db)
	s := services.NewTodoService(repo)
	c := controllers.NewTodoController(s)

	router.HandleFunc("/{id}", api.MakeHandlerFunc(c.GetById)).Methods("GET")
}
