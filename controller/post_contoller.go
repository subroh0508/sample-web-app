package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sample-web-app/model"
)

type PostController interface {
	GetIndex(c *gin.Context)
	GetNew(c *gin.Context)
	PostNew(c *gin.Context)
	GetShow(c *gin.Context)
	GetEdit(c *gin.Context)
	PostEdit(c *gin.Context)
	PostDelete(c *gin.Context)
}

type postController struct {
	db *gorm.DB
}

func NewPostController(db *gorm.DB) PostController {
	return &postController{db}
}

func (controller *postController) GetIndex(c *gin.Context) {
	var posts []model.Post
	controller.db.Find(&posts)
	c.HTML(200, "index.tmpl.html", gin.H{
		"posts": model.MapPostToJson(posts),
	})
}

func (controller *postController) GetNew(c *gin.Context) {
	c.HTML(200, "new.tmpl.html", nil)
}

func (controller *postController) PostNew(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	post := model.Post{Title: title, Content: content}
	controller.db.Create(&post)

	c.Redirect(302, "/")
}

func (controller *postController) GetShow(c *gin.Context) {
	id := c.Param("id")

	var post model.Post
	controller.db.First(&post, id)

	c.HTML(200, "show.tmpl.html", gin.H{
		"post": model.PostToJson(post),
	})
}

func (controller *postController) GetEdit(c *gin.Context) {
	id := c.Param("id")

	var post model.Post
	controller.db.First(&post, id)

	c.HTML(200, "edit.tmpl.html", gin.H{
		"post": model.PostToJson(post),
	})
}

func (controller *postController) PostEdit(c *gin.Context) {
	id := c.Param("id")
	title := c.PostForm("title")
	content := c.PostForm("content")

	var post model.Post
	controller.db.First(&post, id)
	post.Title = title
	post.Content = content
	controller.db.Save(&post)

	c.Redirect(302, "/")
}

func (controller *postController) PostDelete(c *gin.Context) {
	id := c.Param("id")

	controller.db.Delete(&model.Post{}, id)
	c.Redirect(302, "/")
}
