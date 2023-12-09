package example

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(r *gin.Engine) {
	// default: 32MiB
	r.MaxMultipartMemory = 8 << 20 // 8MiB
	r.POST("/25", func(c *gin.Context) {
		file, _ := c.FormFile("file")

		log.Println(file.Filename)
		c.SaveUploadedFile(file, "./"+file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("%s uploaded file!", file.Filename))
	})

	r.POST("/25/multi", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			c.SaveUploadedFile(file, "./"+file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d uploaded file!", len(files)))
	})
}

// curl -X POST http://localhost:8085 \
// -H "Content-Type: multipart/form-data"
