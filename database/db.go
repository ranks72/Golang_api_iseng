package database

import (
	"MALIKI-KARIM/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "Viking_72"
	dBport   = "5432"
	dBname   = "retail_bank"
	db       *gorm.DB
	err      error
)

func StartDb() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dBname, dBport)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err.Error())
	}

	if err != nil {
		log.Fatal("error while tyring to ping the database connection", err.Error())
	}

	fmt.Println("successfully connected to my database")

	//untuk buat tabel baru
	db.Debug().AutoMigrate(entity.User{})
}

func GetDb() *gorm.DB {
	return db
}
