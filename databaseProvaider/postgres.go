package databaseProvaider

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres() {
	dsn := "host=localhost user=postgres password=snegowik2015 dbname=archive port=1234 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to Postgres")
	}

	DB = db
}
