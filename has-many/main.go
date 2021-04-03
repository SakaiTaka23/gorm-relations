package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User has many CreditCards, UserID is the foreign key
type User struct {
	ID          int
	Name        string
	CreditCards []CreditCard
}

type CreditCard struct {
	ID     int
	Number string
	UserID uint
}

func create(db *gorm.DB) {
	user := User{
		Name: "some name",
		CreditCards: []CreditCard{
			{
				Number: "111-111-111-111",
			},
			{
				Number: "222-222-222-222",
			},
			{
				Number: "333-333-333-333",
			},
		},
	}
	db.Create(&user).Save(&user)
}

func read(db *gorm.DB) {
	var user []User
	db.Preload("CreditCards").Find(&user)
	fmt.Println(user)
}

func update(db *gorm.DB) {
	db.Model(&User{}).Where("name = ?", "some name").First(&User{}).Update("name", "changed name")
}

func delete(db *gorm.DB) {
	db.Where("name = ?", "changed name").Delete(&User{})
}

func main() {
	db, _ := gorm.Open(sqlite.Open("has-many.db"), &gorm.Config{})
	db.AutoMigrate(&User{}, &CreditCard{})
	//db = db.Debug()

	create(db)
	read(db)
	update(db)
	read(db)
	delete(db)
	read(db)

	fmt.Println("Done")
}
