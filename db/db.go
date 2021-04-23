package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kougakusai/clubhook-api/models"
)

func Migrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.Event{})
	db.AutoMigrate(&models.Calender{})
	db.AutoMigrate(&models.Vote{})
	db.AutoMigrate(&models.Option{})
	db.AutoMigrate(&models.Cast{})
	return db
}

func Connect() *gorm.DB {
	dburl := os.Getenv("DATABASE_URL")
	if dburl == "" {
		dburl = "host=postgres port=5432 user=example_user dbname=example_db password=password sslmode=disable"
	}
	db, err := gorm.Open(postgres.Open(dburl), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func Close(db *gorm.DB) {
	sqldb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqldb.Close()
}
