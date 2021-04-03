package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// `User` belongs to `Company`, `CompanyID` is the foreign key
type User struct {
	ID        int
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

func create(db *gorm.DB) {
	user := User{
		Name: "my name",
		Company: Company{
			Name: "some company",
		},
	}
	db.Create(&user)
	db.Save(&user)
}

func read(db *gorm.DB) {
	var user User
	// db.Model(&user).Association("Company")
	db.Preload("Company").Find(&user)
	fmt.Println(user)
}

func update(db *gorm.DB) {
	db.Model(&User{}).Where("name = ?", "my name").Update("name", "changed name")
}

func delete(db *gorm.DB) {
	var user User
	db.Where("name = ?", "changed name").Delete(&user)
}

func readRelation(db *gorm.DB) {
	var user User
	var company Company
	db.Preload("Company").Model(&user).First(&user)
	db.Model(&user).Association("Company").Find(&company)
	fmt.Println(user, company)
}

func main() {
	db, _ := gorm.Open(sqlite.Open("belongs-to.db"), &gorm.Config{})
	db.AutoMigrate(&User{}, &Company{})
	//db = db.Debug()

	create(db)
	read(db)
	update(db)
	read(db)
	readRelation(db)
	delete(db)
	read(db)

	fmt.Println("Done")
}
