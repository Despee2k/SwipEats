package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/SwipEats/SwipEats/server/internal/constants"
	"github.com/SwipEats/SwipEats/server/internal/models" 
)

var Conn *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		constants.DB_USER,
		constants.DB_PASSWORD,
		constants.DB_HOST,
		constants.DB_PORT,
		constants.DB_NAME,
	)

	_conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		os.Exit(1)
	}

	Conn = _conn
	fmt.Println("Database connection established successfully")
}

func MigrateModels() {
	if Conn == nil {
		log.Fatal("Database connection is not established")
	}

	err := Conn.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.GroupMembership{},
		&models.Restaurant{},
		&models.Swipe{},
		&models.Match{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	fmt.Println("Database models migrated successfully")
}