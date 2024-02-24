package database

import (
	"fmt"
	"synapsis/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDatabaseConnection(config config.Config) {
	var err error
	dsn := GetDsn(config)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection estabilished")
	}

}

func GetDsn(config config.Config) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.DBUsername, config.DBPassword, config.DBName, config.DBHost, config.DBPort)
}
