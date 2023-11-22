package db

import (
	"fmt"
	"log"
	"os"
	"time"

	lo "github.com/gothew/l-og"
	"github.com/karchx/realword-nx/model"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type OptionsConnection struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

func New(url *OptionsConnection) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", url.Host, url.Port, url.User, url.Password, url.Dbname)

	db, err := gorm.Open(postgres.Open(psqlUrl), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		lo.Errorf("storage err: %v", err)
	} else {
		lo.Infof("DB: %s UP", url.Dbname)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Follow{},
	)
}
