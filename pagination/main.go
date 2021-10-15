package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Article struct {
	ID    uint
	Title string
}

func createFakeData(db *gorm.DB) {
	var articles []Article
	i := 1
	for i <= 50 {
		article := Article{
			Title: fmt.Sprintf("article%d", i),
		}
		articles = append(articles, article)
		i++
	}
	db.Create(&articles)
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func main() {
	db, _ := gorm.Open(sqlite.Open("pagination.db"), &gorm.Config{})
	db.AutoMigrate(&Article{})
	// db.Debug()

	// createFakeData(db)

	var article []Article
	db.Scopes(Paginate(1, 10)).Find(&article)
	fmt.Println(article)
}
