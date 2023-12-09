package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type FormA struct {
	Foo string `json:"foo" binding:"required"`
}

type FormB struct {
	Bar string `json:"bar" binding:"required"`
}

/*
- c.ShouldBindWith lưu body vào context trước khi bind
điều nó có ảnh hưởng nhẹ đến performance => không nên sử dụng nếu chỉ gọi binding 1 lần
- c.ShouldBindWith chỉ cần thiết cho 1 số format: JSON, XML, MsgPack, ProtoBuf. 
Đối với các format khác: Query, Form, FormPost, FormMultipart có thể được gọi bằng c.ShouldBind() nhiều lần mà không ảnh hưởng đến hiệu suất
*/

func BindBodyIntoDifferentStruct(r *gin.Engine) {
	r.POST("/24", func(c *gin.Context) {
		objA := FormA{}
		objB := FormB{}
		// c.ShouldBind sử dụng c.Request.Body và nó không thể tái sử dụng.
		if errA := c.ShouldBind(&objA); errA == nil {
			c.String(http.StatusOK, `the body should be formA`)

			// Luôn luôn lỗi bởi vì c.Request.Body nhận EOF
		} else if errB := c.ShouldBind(&objB); errB == nil {
			c.String(http.StatusOK, `the body should be formB`)
		} else {
			c.String(http.StatusBadRequest, `Not valid`)
		}
	})

	r.POST("/24/valid", func(c *gin.Context) {
		objA := FormA{}
		objB := FormB{}
		// Đọc c.Request.Body và lưu trữ nó vào context
		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			c.String(http.StatusOK, `the body should be formA`)

			// Tại thời điểm này, nó sử dụng lại body được lưu trong context
		} else if errB := c.ShouldBindWith(&objB, binding.JSON); errB == nil {
			c.String(http.StatusOK, `the body should be formB`)
		} else {
			c.String(http.StatusBadRequest, `Not valid`)
		}
	})
}
