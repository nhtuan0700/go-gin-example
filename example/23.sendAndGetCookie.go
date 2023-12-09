package example

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SendAndGetCookie(r *gin.Engine) {
	r.GET("/23", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")

		if err != nil {
			cookie = "NotSet"

			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s\n", cookie)
	})
}
