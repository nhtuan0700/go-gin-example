package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(r *gin.Engine) {
	r.GET("/19", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
	})

	r.POST("/19/post", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/19/test")
	})

	r.GET("/19/test", func(c *gin.Context) {
		c.Request.URL.Path = "/19/test2"
		r.HandleContext(c)
	})
	r.GET("/19/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})
}
