package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sample-web-app/model"
)

type BlogController interface {
	Index(c *gin.Context)
	// Create(c *gin.Context)
	// Show(id uint, c *gin.Context)
	// Update(id uint, c *gin.Context)
	// Delete(id uint, c *gin.Context)
}

type blogController struct {
	db *gorm.DB
}

func NewBlogController(db *gorm.DB) BlogController {
	return &blogController{db}
}

func (controller *blogController) Index(c *gin.Context) {
	var blogs []model.Blog
	controller.db.Find(&blogs)
	c.HTML(200, "index.html", gin.H{
		"blogs": blogs,
	})
}
