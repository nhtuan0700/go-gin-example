package example

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonP(r *gin.Engine) {
	r.GET("/JSONP", func(c *gin.Context) {
		callback := c.Query("callback")
		log.Println(callback)
		data := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(http.StatusOK, data)
	})
}

// curl "localhost:8088/JSONP?callback=x"
