package model

import (
	"gorm.io/gorm"
	"strconv"
)

type Post struct {
	gorm.Model
	Title   string
	Content string
}

func (p Post) ShowHref() string {
	return "/posts/" + strconv.FormatUint(uint64(p.ID), 10)
}

func (p Post) EditHref() string {
	return "/posts/" + strconv.FormatUint(uint64(p.ID), 10) + "/edit"
}

func (p Post) DeleteHref() string {
	return "/posts/" + strconv.FormatUint(uint64(p.ID), 10) + "/delete"
}

func MapPostToJson(posts []Post) []map[string]string {
	result := make([]map[string]string, 0, len(posts))

	for _, post := range posts {
		result = append(result, PostToJson(post))
	}

	return result
}

func PostToJson(post Post) map[string]string {
	return map[string]string{
		"id":             strconv.FormatUint(uint64(post.ID), 10),
		"title":          post.Title,
		"content":        post.Content,
		"get_show_href":  post.ShowHref(),
		"get_edit_href":  post.EditHref(),
		"post_edit_href": post.ShowHref(),
		"delete_href":    post.DeleteHref(),
	}
}
