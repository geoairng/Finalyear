package database

import (
	"fmt"
	"log"

	"github.com/geoairng/Finalyear/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUp() (*gorm.DB, error) {

	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	var (
		db  *gorm.DB
		err error
	)

	var dsn string

	dsn = fmt.Sprintf("host=localhost user=postgres password=password123 dbname=finalyear port=5432 sslmode=disable TimeZone=Africa/Lagos")

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(&model.User{})

	if err != nil {
		log.Println(err)
	}

	return db, err
}
