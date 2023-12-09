package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SecureJSON(r *gin.Engine) {
	r.GET("/21", func (c *gin.Context)  {
		name := []string{"lena", "austin", "foo"}
		r.SecureJsonPrefix(")]}',\n")
		c.SecureJSON(http.StatusOK, name)
	})
}
