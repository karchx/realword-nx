package postgres

import (
	"database/sql"
	"fmt"

	log "github.com/gothew/l-og"
	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type UrlDB struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func Open(url *UrlDB) (*DB, error) {
	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", url.Host, url.Port, url.User, url.Password, url.Dbname)
	db, err := sql.Open("postgres", psqlUrl)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Info("successfully connected to database")
	return &DB{db}, nil
}
