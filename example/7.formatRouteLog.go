package example

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FormatRouteLog(r *gin.Engine) {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.POST("/foo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "foo")
	})
}
