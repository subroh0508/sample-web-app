package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sample-web-app/controller"
	"sample-web-app/model"
)

func main() {
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Blog{})

	var blog model.Blog
	db.First(&blog, 1)
	if blog.ID == 0 {
		db.Create(&model.Blog{Title: "First Post", Content: "Hello, World!"})
	}

	blogController := controller.NewBlogController(db)

	r := gin.Default()

	r.LoadHTMLGlob("views/*")
	r.GET("/", blogController.Index)

	r.Run(":8080")
}
