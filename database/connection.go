package database

import (
	"e-commerce-golang/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

var DB *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=e_commerce sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Comment{})
	return db
}

func GetDB() *gorm.DB {
	if DB == nil {
		DB = Init()
		var sleep = time.Duration(1)
		for DB == nil {
			sleep = sleep * 2
			fmt.Printf("Database in unavalible. Wait for %d sec.\n", sleep)
			time.Sleep(sleep * time.Second)
			DB = Init()
		}
	}
	return DB
}
