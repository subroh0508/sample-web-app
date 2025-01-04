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
	return "/posts/" + strconv.FormatUint(uint64(p.ID), 10)
}

func MapPostToJson(posts []Post) []map[string]string {
	result := make([]map[string]string, 0, len(posts))

	for _, post := range posts {
		result = append(result, map[string]string{
			"id":          strconv.FormatUint(uint64(post.ID), 10),
			"title":       post.Title,
			"content":     post.Content,
			"show_href":   post.ShowHref(),
			"edit_href":   post.EditHref(),
			"delete_href": post.DeleteHref(),
		})
	}

	return result
}
