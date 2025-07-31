package config

import (
	"log"
	"vision-service/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/vision-db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	err = db.AutoMigrate(&entity.Post{})
	if err != nil {
		log.Fatal("Gagal migrate database:", err)
	}

	log.Println("Database terkoneksi dan migrasi berhasil.")
	return db
}
