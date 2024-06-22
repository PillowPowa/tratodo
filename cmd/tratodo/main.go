package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"tratodo/internal/api/routes"
	"tratodo/internal/config"
	"tratodo/internal/storage"

	_ "tratodo/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	Env *config.Config
	DB  *sql.DB
}

// @title Tratodo API
// @version 0.1
// @description This is a simple todo list API

// @BasePath /

// @schemes http
// @produce json
// @consumes json

// @SecurityDefinitions jwt
// @type jwt
// @in cookie

func main() {
	app := App{}

	app.Env = config.MustLoad()

	db, err := storage.NewSqlite(app.Env.StoragePath)
	if err != nil {
		panic(err)
	}
	app.DB = db

	addr := fmt.Sprintf("%s:%v", app.Env.HTTP.Host, app.Env.HTTP.Port)
	swaggerUrl := fmt.Sprintf("http://%s/docs/doc.json", addr)
	router := mux.NewRouter()

	router.PathPrefix("/docs/").Handler(
		httpSwagger.Handler(httpSwagger.URL(swaggerUrl)),
	)

	todoGroup := router.PathPrefix("/todo").Subrouter()
	routes.NewTodoRoute(app.DB, todoGroup)

	authGroup := router.PathPrefix("/auth").Subrouter()
	routes.NewAuthRoute(app.DB, &app.Env.JWT, authGroup)

	http.Handle("/", router)

	log.Printf("API ready recieve ur requests on addr: http://%s", addr)
	srv := http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  time.Duration(app.Env.HTTP.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(app.Env.HTTP.WriteTimeout) * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
