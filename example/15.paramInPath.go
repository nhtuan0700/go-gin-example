package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParamInPath(r *gin.Engine) {
	// match: /15/tuan
	r.GET("/15/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// match: /15/tuan/, /15/tuan/send
	r.GET("/15/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		message := name + " is " + action

		c.String(http.StatusOK, message)
	})
}
