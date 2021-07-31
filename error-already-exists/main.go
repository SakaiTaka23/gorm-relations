package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string `gorm:"unique"`
}

func create(db *gorm.DB) {
	user := User{
		Name: "some name",
	}
	if err := db.Create(&user).Save(&user).Error; errors.Is(err, gorm.ErrInvalidTransaction) {
		fmt.Printf("error caught %s\n", err)
	}
}

func main() {
	db, _ := gorm.Open(sqlite.Open("select.db"), &gorm.Config{})
	db.AutoMigrate(&User{})

	create(db)
	create(db)

	fmt.Println("Done")
}
