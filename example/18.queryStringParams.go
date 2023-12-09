package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryStringParams(r *gin.Engine) {
	r.GET("18", func (c *gin.Context)  {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
}
