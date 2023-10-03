package main

import (
	"github.com/karchx/realword-nx/db"
	"github.com/karchx/realword-nx/handler"
	"github.com/karchx/realword-nx/router"
	"github.com/karchx/realword-nx/store"

	log "github.com/gothew/l-og"
)

func main() {
	r := router.New()

	d := db.New(&db.OptionsConnection{
		Host:     "0.0.0.0",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Dbname:   "realworld",
	})

	db.AutoMigrate(d)
	us := store.NewUserStore(d)

	h := handler.NewHandler(us)
	h.Register(r)

	err := r.Listen(":3000")
	if err != nil {
		log.Errorf("%v", err)
	}
}
