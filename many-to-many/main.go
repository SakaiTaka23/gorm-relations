package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	ID       uint
	LangName string
}

func create(db *gorm.DB) {
	user := User{
		Name: "some name",
		Languages: []Language{
			{
				LangName: "PHP",
			},
			{
				LangName: "Go",
			},
		},
	}
	db.Create(&user).Save(&user)
}

func read(db *gorm.DB) {
	var user []User
	db.Preload("Languages").Find(&user)
	fmt.Println(user)
}

func update(db *gorm.DB) {
	db.Model(&User{}).Where("name = ?", "some name").Update("name", "changed name")
}

func delete(db *gorm.DB) {
	var user []User
	db.Preload("Languages").Where("name = ?", "changed name").Find(&user)
	db.Model(&user).Association("Languages").Clear()
	db.Delete(&user)
}

func main() {
	db, _ := gorm.Open(sqlite.Open("many-to-many.db"), &gorm.Config{})
	db.AutoMigrate(&User{}, &Language{})
	// db = db.Debug()

	create(db)
	read(db)
	update(db)
	read(db)
	delete(db)
	read(db)

	fmt.Println("Done")
}
