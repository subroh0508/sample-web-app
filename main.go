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

	err = db.AutoMigrate(&model.Post{})
	if err != nil {
		panic("failed to migrate database")
	}

	postController := controller.NewPostController(db)

	r := gin.Default()

	r.LoadHTMLGlob("resources/*")
	r.GET("/", postController.GetIndex)
	r.GET("/posts/new", postController.GetNew)
	r.POST("/posts/new", postController.PostNew)
	r.GET("/posts/:id", postController.GetShow)
	r.GET("/posts/:id/edit", postController.GetEdit)
	r.POST("/posts/:id", postController.PostEdit)
	r.POST("/posts/:id/delete", postController.PostDelete)

	err = r.Run(":8080")
	if err != nil {
		panic("failed to run server")
	}
}
