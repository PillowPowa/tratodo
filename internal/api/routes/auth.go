package routes

import (
	"database/sql"
	"time"
	"tratodo/internal/api/controllers"
	"tratodo/internal/api/middlewares"
	"tratodo/internal/config"
	"tratodo/internal/repository"
	"tratodo/internal/services"
	"tratodo/pkg/api"

	"github.com/gorilla/mux"
)

func NewAuthRoute(db *sql.DB, env *config.JWTConfig, router *mux.Router) {
	repo := repository.NewUserRepository(db)
	s := services.NewAuthService(repo, time.Duration(env.ExpDays)*time.Hour*24)
	c := controllers.NewAuthController(s)

	router.HandleFunc("/login", api.MakeHandlerFunc(c.Login)).Methods("POST")
	router.HandleFunc("/register", api.MakeHandlerFunc(c.Register)).Methods("POST")
	router.HandleFunc("/logout", api.MakeHandlerFunc(middlewares.WithAuth(c.Logout))).Methods("POST")
}
