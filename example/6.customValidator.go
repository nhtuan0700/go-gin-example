package example

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Booking contains binded and validated data.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func CustomValidator(r *gin.Engine) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	// curl "localhost:8088/6/bookable?check_in=2118-04-16&check_out=2118-04-17"
	// {"message":"Booking dates are valid!"}
	// curl "localhost:8088/6/bookable?check_in=2118-03-10&check_out=2118-03-09"
	// {"error":"Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"}
	r.GET("/6/1", getBookable)
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// Struct validation level
type User struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `binding:"required,email"`
}

func UserStructLevelValidation(sl validator.StructLevel) {
	// user := structLevel.CurrentStruct.Interface().(User)
	user := sl.Current().Interface().(User)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		sl.ReportError(user.FirstName, "FirstName", "fname", "fnameorlname", "")
		sl.ReportError(user.LastName, "LastName", "lname", "fnameorlname", "")
	}
}

func CustomValidator2(r *gin.Engine) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(UserStructLevelValidation, User{})
	}
	/*
		curl -s -X POST http://localhost:8088/6/2 \
		-H 'content-type: application/json' \
		-d '{}'
		{
			"error": "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'User.FirstName' Error:Field validation for 'FirstName' failed on the 'fnameorlname' tag\nKey: 'User.LastName' Error:Field validation for 'LastName' failed on the 'fnameorlname' tag",
			"message": "User validation failed!"
		}
		curl -X POST http://localhost:8085/user \
		-H 'content-type: application/json' \
		-d '{"fname": "George", "lname": "Costanza", "email": "george@vandaley.com"}'
		{"message":"User validation successful."}
	*/

	r.POST("/6/2", validateUser)
}

func validateUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User validation successful."})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User validation failed!",
			"error":   err.Error(),
		})
	}
}
