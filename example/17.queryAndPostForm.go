package example

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func QueryAndPostForm(r *gin.Engine) {
	r.POST("/17", func (c *gin.Context)  {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page %s; name: %s; message: %s", id, page, name, message)
	})
}

// curl http://localhost:8088/17?id=123&page=1 \
// -H 'Content-Type:application/x-www-form-urlencoded' \
// -X POST \
// -d "name=tuan&message=this_is_great"
