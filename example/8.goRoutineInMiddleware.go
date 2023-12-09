package example

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func GoRoutineInMiddleware(r *gin.Engine) {
	r.GET("/8/long-async", func(c *gin.Context) {
		cCp := c.Copy()
		// When starting new Goroutines inside a middleware or handler, you SHOULD NOT use the original context inside it, you have to use a read-only copy.
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done in path " + cCp.Request.URL.Path)
		}()
	})
	r.GET("/8/long-sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done in path " + c.Request.URL.Path)
	})
}
