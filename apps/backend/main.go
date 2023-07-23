package main

import (
	log "github.com/gothew/l-og"
	"github.com/karchx/realword-nx/conduit"
	"github.com/karchx/realword-nx/postgres"
	"github.com/karchx/realword-nx/server"
)

type config struct {
	port  string
	dbURI postgres.UrlDB
}

func main() {
	cfg := envConfig()

	db, err := postgres.Open(&cfg.dbURI)

	if err != nil {
		log.Fatalf("cannot open database: %v", err)
	}
	db.AutoMigrate(&conduit.User{})
	log.Info("Migrated")

	srv := server.NewServer(db)
	log.Fatal(srv.Run(cfg.port))
}

func envConfig() config {
	/*
		port, _ := os.LookupEnv("PORT")
			if !ok {
				panic("PORT not provided")
			}

			dbURI, ok := os.LookupEnv("POSTGRESQL_URL")

			if !ok {
				panic("POSTGRESQL_URL not provided")
			}*/

	return config{port: "5001", dbURI: postgres.UrlDB{
		Host:     "0.0.0.0",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Dbname:   "realword",
	}}
}
