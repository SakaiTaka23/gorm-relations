package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
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

func getFind(db *gorm.DB) {
	var user []User
	result := db.Where("name = ?", "name3").Find(&user)
	// エラーは検知されない
	fmt.Println(result.Error)
	fmt.Println(user)
}

func getFirst(db *gorm.DB) {
	var User []User
	if err := db.Where("name = ?", "name3").First(&User).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("error detected! ErrRecordNotFound")
	}
	// 検知可能 コード自体は問題なく動く
	// error : record not found
	fmt.Println(User)
}

func getLast(db *gorm.DB) {
	var User []User
	if err := db.Where("name = ?", "name3").Last(&User).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("error detected! ErrRecordNotFound")
	}
	// 検知可能 コード自体は問題なく動く
	// error : record not found
	fmt.Println(User)
}

func getTake(db *gorm.DB) {
	var User []User
	if err := db.Where("name = ?", "name3").Take(&User).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("error detected! ErrRecordNotFound")
	}
	fmt.Println(User)
}

func main() {
	db, _ := gorm.Open(sqlite.Open("select.db"), &gorm.Config{})
	db.AutoMigrate(&User{})
	// db.Debug()

	createFakeData(db)

	getFind(db)
	getFirst(db)
	getLast(db)
	getTake(db)

	fmt.Println("Done")
}
