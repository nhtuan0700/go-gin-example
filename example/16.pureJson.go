package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PureJSON(r *gin.Engine) {
	r.GET("/16", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello world</b>",
		})
	})

	// not available: 1.6 and lower
	r.GET("/16/encode", func (c *gin.Context)  {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello world</b>",
		})
	})
}
