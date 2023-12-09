package example

import (
	"log"

	"github.com/gin-gonic/gin"
)

func MapQueryStringPostOrFormParams(r *gin.Engine) {
	r.POST("/12", func(c *gin.Context) {
		log.Println(c.Query("abc"))
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		log.Printf("ids: %v; names: %v", ids, names)
	})
}
// curl --globoff http://localhost:8088/12?ids[a]=1234&ids[b]=hello \
// -H 'Content-Type:application/x-www-form-urlencoded' \
// -X POST \
// -d "names[first]=thinkerou&names[second]=tianou"
