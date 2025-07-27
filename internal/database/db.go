// Package db
package database

import (
	"fmt"
	"log"

	"github.com/Guglahai/tasks-service/internal/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(creds *configs.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		creds.Host, creds.User, creds.Password, creds.Database, creds.Port, creds.SSLMode)
	var err error

	fmt.Println(dsn)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to database %s", err)
	}

	return db, nil
}
