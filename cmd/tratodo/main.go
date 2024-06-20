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

	"github.com/gorilla/mux"
)

type App struct {
	Env *config.Config
	DB  *sql.DB
}

func main() {
	app := App{}

	app.Env = config.MustLoad()

	db, err := storage.NewSqlite(app.Env.StoragePath)
	if err != nil {
		panic(err)
	}
	app.DB = db

	router := mux.NewRouter()

	todoGroup := router.PathPrefix("/todo").Subrouter()
	routes.NewTodoRoute(db, todoGroup)

	authGroup := router.PathPrefix("/auth").Subrouter()
	routes.NewAuthRoute(db, &app.Env.JWT, authGroup)

	http.Handle("/", router)

	addr := fmt.Sprintf("%s:%v", app.Env.HTTP.Host, app.Env.HTTP.Port)
	log.Printf("API ready recieve ur requests on addr: http://%s", addr)
	srv := http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  time.Duration(app.Env.HTTP.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(app.Env.HTTP.WriteTimeout) * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
