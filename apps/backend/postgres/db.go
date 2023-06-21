package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	log "github.com/gothew/l-og"
	_ "github.com/lib/pq"
)

type DB struct {
	*gorm.DB
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
	db, err := gorm.Open(postgres.Open(psqlUrl), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Info("successfully connected to database")
	return &DB{db}, nil
}
