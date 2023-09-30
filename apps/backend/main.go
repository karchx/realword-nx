package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karchx/realword-nx/db"

	log "github.com/gothew/l-og"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	d := db.New(&db.OptionsConnection{
		Host:     "0.0.0.0",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Dbname:   "realworld",
	})

	db.AutoMigrate(d)

	log.Fatal(app.Listen(":3000"))
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
