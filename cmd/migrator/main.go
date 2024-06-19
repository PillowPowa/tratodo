package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var driver string
	flag.StringVar(&driver, "driver", "sqlite", "database driver")

	var isDownMigration bool
	flag.BoolVar(&isDownMigration, "down", false, "up migration")

	var m *migrate.Migrate
	switch driver {
	case "sqlite":
		m = makeSqliteMigration()
	default:
		panic("unknown driver")
	}

	if isDownMigration {
		if err := m.Down(); err != nil {
			panic(err)
		}
	} else {
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to apply")

				return
			}

			panic(err)
		}
	}
}

func makeSqliteMigration() *migrate.Migrate {
	var dataPath, migrationsPath, migrationsTable string

	flag.StringVar(&dataPath, "data", "", "path to database file")
	flag.StringVar(&migrationsPath, "migrations", "", "path to migrations directory")
	flag.StringVar(&migrationsTable, "table", "migrations", "migrations table name")
	flag.Parse()

	if dataPath == "" {
		panic("data path is empty")
	}

	if migrationsPath == "" {
		panic("migrations path is empty")
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", dataPath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}
	return m
}
