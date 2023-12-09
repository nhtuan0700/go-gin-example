package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GroupRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/a", func(c *gin.Context) {
			c.String(http.StatusOK, "Message: a1")
		})
		v1.GET("/b", func(c *gin.Context) {
			c.String(http.StatusOK, "Message: b1")
		})
	}
	v2 := r.Group("/v2")

	v2.GET("/a", func(c *gin.Context) {
		c.String(http.StatusOK, "Message: a2")
	})
	v2.GET("/b", func(c *gin.Context) {
		c.String(http.StatusOK, "Message: b2")
	})
}
