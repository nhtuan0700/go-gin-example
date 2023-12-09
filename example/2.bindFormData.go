package example

import "github.com/gin-gonic/gin"

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

func BindFormData(r *gin.Engine) {
	r.GET("/2/getb", GetDataB)
	r.GET("/2/getc", GetDataC)
	r.GET("/2/getd", GetDataD)
}

// curl "http://localhost:8088/2/getb?field_a=hello&field_b=world"
// {"a":{"FieldA":"hello"},"b":"world"}
// curl "http://localhost:8088/2/getc?field_a=hello&field_c=world"
// {"a":{"FieldA":"hello"},"c":"world"}
// curl "http://localhost:8088/2/getd?field_x=hello&field_d=world"
// {"d":"world","x":{"FieldX":"hello"}}
