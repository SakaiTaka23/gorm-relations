package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

type APIUser struct {
	Name string
	Age  int
}

func createFakeData(db *gorm.DB) {
	users := []User{
		{
			Name: "name1",
			Age:  20,
		},
		{
			Name: "name2",
			Age:  30,
		},
	}
	db.Create(&users)
}

func selectMethod(db *gorm.DB) {
	var user []User
	db.Select("Name", "Age").Find(&user)
	fmt.Println(user)
}

func smartSelect(db *gorm.DB) {
	var apiUser []APIUser
	db.Model(&User{}).Find(&apiUser)
	fmt.Println(apiUser)
}

func main() {
	db, _ := gorm.Open(sqlite.Open("select.db"), &gorm.Config{})
	db.AutoMigrate(&User{})
	// db.Debug()

	//createFakeData(db)

	selectMethod(db)
	smartSelect(db)

	fmt.Println("Done")
}
