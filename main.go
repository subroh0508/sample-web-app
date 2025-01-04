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

	db.AutoMigrate(&model.Post{})

	var post model.Post
	db.First(&post, 1)
	if post.ID == 0 {
		db.Create(&model.Post{Title: "First Post", Content: "Hello, World!"})
	}

	postController := controller.NewPostController(db)

	r := gin.Default()

	r.LoadHTMLGlob("views/*")
	r.GET("/", postController.Index)

	r.Run(":8080")
}
