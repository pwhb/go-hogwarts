package database

import (
	"log"

	"github.com/pwhb/go-gin-hogwarts/pkg/models/house"
	"github.com/pwhb/go-gin-hogwarts/pkg/models/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		panic("Failed to connect database")
	} else {
		log.Println("Database connected successfully!")
	}

	if err = DB.AutoMigrate(&user.User{}); err != nil {
		log.Println(err)
	} else {
		log.Println("User Model migrated successfully!")
	}

	if err = DB.AutoMigrate(&house.House{}); err != nil {
		log.Println(err)
	} else {
		log.Println("House Model migrated successfully!")
	}

	// if err = DB.AutoMigrate(&employee.Employee{}); err != nil {
	// 	log.Println(err)
	// } else {
	// 	log.Println("Employee Model migrated successfully!")
	// }

}
