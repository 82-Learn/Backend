package database

import (
	"fmt"

	"82learn.com/backend/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	DBConn *gorm.DB
)

func ConnectDB() {
	var err error

	db, err := gorm.Open("postgres", "user=postgres dbname="" host="" port="" password=""")
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected to ReactDB")

	db.AutoMigrate(&models.Customer{})

	fmt.Println("Database Migrated")

}
