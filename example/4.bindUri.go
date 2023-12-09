package example

import "github.com/gin-gonic/gin"

type PersonURI struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func BindURI(r *gin.Engine) {
	r.GET("/4/:name/:id", func(c *gin.Context) {
		var person PersonURI
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
}

// curl localhost:8088/4/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
// curl localhost:8088/4/thinkerou/not-uuid
