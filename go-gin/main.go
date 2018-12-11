package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("hello %s", c.Param("name")),
		})
	})

	type ParamList struct {
		Name string `form:"name" binding:"required"`
		Age  string `form:"age" binding:"required"`
	}

	r.GET("/param", func(c *gin.Context) {
		var param ParamList
		if err := c.ShouldBindQuery(&param); err != nil {
			fmt.Println(err)
			// log.Fatal(err)
			return
		}
		fmt.Println(param)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("name: %s , age: %s", param.Name, param.Age),
		})
	})

	r.Run(":8888") // listen and serve on 0.0.0.0:8080
}
