package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindAsciiJson(router *gin.Engine) {
	router.GET("/1", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})
}
