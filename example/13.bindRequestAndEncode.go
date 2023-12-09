package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

func BindRequestAndEncode(r *gin.Engine) {
	// For JSON
	r.POST("/13", func(c *gin.Context) {
		var jsonRequest LoginRequest
		if err := c.ShouldBindJSON(&jsonRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if jsonRequest.Email != "tuan@test.com" || jsonRequest.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// For form data
	r.POST("/13/form", func(c *gin.Context) {
		var formRequest LoginRequest

		if err := c.ShouldBind(&formRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if formRequest.Email != "tuan@test.com" || formRequest.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.POST("13/encode", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
}

// curl http://localhost:8088/13 \
// -X POST \
// -d '{"email": "tuan@test.com", "password":""}'

// curl http://localhost:8088/13/form \
// -H "Content-Type: application/x-www-form-urlencoded" \
// -d "email=tuan@test.com&password=123"

// curl --form email=tuan@test.com --form password=123 http://localhost:8088/13/form
