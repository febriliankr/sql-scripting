package main

import (
	"log"

	"github.com/febriliankr/go-sql-scripting/adapter"
	"github.com/febriliankr/go-sql-scripting/config"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.InitConfig()
	db, err := adapter.NewDB(cfg)
	if err != nil {
		panic(err)
	}

	app := NewApp(db)

	log.Println("app", app)

}

type App struct {
	Db *sqlx.DB
}

func NewApp(db *sqlx.DB) *App {
	return &App{
		Db: db,
	}
}
