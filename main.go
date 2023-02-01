package main

import (
	"github.com/febriliankr/go-sql-scripting/adapter"
	"github.com/febriliankr/go-sql-scripting/app"
	"github.com/febriliankr/go-sql-scripting/config"
)

func main() {
	cfg := config.InitConfig()
	db, err := adapter.NewDB(cfg)
	if err != nil {
		panic(err)
	}

	app := app.NewOpenexam(db)

	err = app.ImportAll()

	if err != nil {
		panic(err)
	}

}
