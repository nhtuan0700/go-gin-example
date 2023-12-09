package example

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Person14 struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func OnlyBindQueryIgnoreForm(r *gin.Engine) {
	r.Any("14", func(c *gin.Context) {
		var person Person
		if c.ShouldBindQuery(&person) == nil {
			log.Println("====== Only Bind By Query String ======")
			log.Println(person.Name)
			log.Println(person.Address)
		}
		c.String(200, "Success")
	})
}
