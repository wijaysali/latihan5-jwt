package config

import (
	"latihan5-jwt/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=userr18 password=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Tidak berhasil koneksi database", err)
	}

	db.AutoMigrate(&models.User{})
	return db
}
