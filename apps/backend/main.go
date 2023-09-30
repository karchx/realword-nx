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

/*type config struct {
	port  string
	dbURI postgres.UrlDB
}

func main() {
	cfg := envConfig()

	db, err := postgres.Open(&cfg.dbURI)

	if err != nil {
		log.Fatalf("cannot open database: %v", err)
	}
	db.AutoMigrate(&conduit.User{}, &conduit.Follow{})
	log.Info("Migrated")

	srv := server.NewServer(db)
	log.Fatal(srv.Run(cfg.port))
}

func envConfig() config {
	return config{port: "5001", dbURI: postgres.UrlDB{
		Host:     "0.0.0.0",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Dbname:   "realworld",
	}}
}*/
