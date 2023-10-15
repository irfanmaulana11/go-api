package dbrepositories

import (
	"fmt"
	"go-api/app/repositories"
	"go-api/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

type DBRepository struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) repositories.DBRepositoryInterface {
	return &DBRepository{db}
}

// create new Postgres database connection with gorm
func NewPostgreSQLConn(conf config.PostgreSQLConfiguration) (*gorm.DB, error) {
	connURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s %s",
		conf.DBHost,
		conf.DBPort,
		conf.DBUser,
		conf.DBName,
		conf.DBPassword,
		conf.DBOptions)

	db, err := gorm.Open("postgres", connURL)
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}
	db.LogMode(true)

	// Auto migrate
	err = goose.Up(db.DB(), "app/repositories/dbrepositories/migrations")
	if err != nil {
		log.Println("Error when running migration:", err.Error())
	}

	return db, nil
}
