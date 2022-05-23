package config

import (
	"fmt"
	"lab-gorm-crud/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Variable global
var DB *gorm.DB

// Database Connection
func ConnectDB() *gorm.DB {
	host := "localhost"
	username := "postgres"
	password := "admin"
	database := "gorm_db"
	port := "5432"
	zone := "Asia/Jakarta"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, username, password, database, port, zone)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})

	if err != nil {
		panic("Can't connect to database")
	}

	// Mapping models
	DB.AutoMigrate(&entity.Products{}, &entity.Orders{}, &entity.OrderDetails{})
	return DB
}
