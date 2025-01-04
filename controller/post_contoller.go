package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sample-web-app/model"
)

type PostController interface {
	Index(c *gin.Context)
	// Create(c *gin.Context)
	// Show(id uint, c *gin.Context)
	// Update(id uint, c *gin.Context)
	// Delete(id uint, c *gin.Context)
}

type postController struct {
	db *gorm.DB
}

func NewPostController(db *gorm.DB) PostController {
	return &postController{db}
}

func (controller *postController) Index(c *gin.Context) {
	var posts []model.Post
	controller.db.Find(&posts)
	c.HTML(200, "index.html", gin.H{
		"blogs": model.MapPostToJson(posts),
	})
}
