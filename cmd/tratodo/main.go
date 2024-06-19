package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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
	http.Handle("/", router)

	addr := fmt.Sprintf("%s:%v", app.Env.HTTP.Host, app.Env.HTTP.Port)
	log.Printf("API ready recieve ur requests on addr: http://%s", addr)
	_ = http.ListenAndServe(addr, router)
}
