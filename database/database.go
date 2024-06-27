package database

import (
	"context"
	"get-on/ent"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	Client *ent.Client
)

func SetupDatabase() {
	var err error
	Client, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
