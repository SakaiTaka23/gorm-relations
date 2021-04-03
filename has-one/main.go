package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User has one CreditCard, CreditCardID is the foreign key
type User struct {
	ID         int
	Name       string
	CreditCard CreditCard
}

type CreditCard struct {
	ID     int
	Number string
	UserID uint
}

func create(db *gorm.DB) {
	user := User{
		Name: "some name",
		CreditCard: CreditCard{
			Number: "111-111-111-111",
		},
	}
	db.Create(&user).Save(&user)
}

func read(db *gorm.DB) {
	var user []User
	db.Preload("CreditCard").Find(&user)
	fmt.Println(user)
}

func update(db *gorm.DB) {
	db.Model(&User{}).Where("name = ?", "some name").First(&User{}).Update("name", "changed name")
}

func delete(db *gorm.DB) {
	db.Where("name = ?", "changed name").Delete(&User{})
}

func main() {
	db, _ := gorm.Open(sqlite.Open("has-one.db"), &gorm.Config{})
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
