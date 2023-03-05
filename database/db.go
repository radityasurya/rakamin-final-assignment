package database

import (
	"fmt"
	"log"

	"github.com/radityasurya/btpn-syariah-final/config"
	"github.com/radityasurya/btpn-syariah-final/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database: ", err)
	}
	fmt.Println("? Connected Successfully to the Database!")
}

func GetDB() *gorm.DB {
	return DB
}

func MigrateDB() {
	fmt.Println("Starting migration...")
	DB.Debug().AutoMigrate(&models.User{})
	fmt.Println("Migration complete!")
}
