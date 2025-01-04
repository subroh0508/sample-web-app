package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")
	r.GET("/", indexHandler)

	r.Run(":8080")
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
