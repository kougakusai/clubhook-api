package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/kougakusaiHPTeam/clubhook-api/models"
)

func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Event{})
	return db
}

func Connect() *gorm.DB {
	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		dburl = "host=postgres port=5432 user=example_user dbname=example_db password=password sslmode=disable"
	}
	db, err := gorm.Open("postgres", dburl)
	if err != nil {
		panic(err)
	}
	return db
}
