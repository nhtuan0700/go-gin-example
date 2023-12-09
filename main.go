package main

import (
	"example/gin-api/example"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	ExampleMain()

	// GraceFulRestartStopMain()
	// RunMultipleServiceMain()
}

func ExampleMain() {
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	// example.FormatRouteLog(router)

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "api is ok"})
	})

	example.BindAsciiJson(router)
	example.BindFormData(router)
	example.BindQueryString(router)
	example.BindURI(router)

	// This example for custom middleware: example/5.customMiddleware
	// router = gin.New()
	// example.CustomMiddleware(router)

	example.CustomValidator(router)
	example.CustomValidator2(router)

	example.GoRoutineInMiddleware(router)

	example.GroupRouter(router)

	example.JsonP(router)
	example.MapQueryStringPostOrFormParams(router)

	example.BindRequestAndEncode(router)
	example.OnlyBindQueryIgnoreForm(router)

	example.ParamInPath(router)
	example.PureJSON(router)

	example.QueryAndPostForm(router)
	example.QueryStringParams(router)

	example.Redirect(router)

	example.SecureJSON(router)

	example.ServingDataFromReader(router)
	example.SendAndGetCookie(router)
	example.BindBodyIntoDifferentStruct(router)
	
	example.UploadFile(router)
	
	example.BasicAuthMiddleware(router)
	router.Run(":8088")
}
