package main

import (
	"database/sql"
	"log"
	"tratodo/internal/config"
	"tratodo/internal/storage"
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
	log.Println("Connected to database")
}
