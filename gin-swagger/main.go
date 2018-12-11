package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "leeif.me/Go-test/gin-swagger/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host

func main() {
	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", APIPing)
	r.GET("/hello/:name", APIHello)
	r.GET("/param", APIParam)

	r.Run(":8888") // listen and serve on 0.0.0.0:8888
}

func (m *MessageInfo) RespMessage() MessageInfo {
	return MessageInfo{
		Message: m.Message,
	}
}

type MessageInfo struct {
	Message string `json:"messge"`
}

type ParamList struct {
	Name string `form:"name" binding:"required"`
	Age  string `form:"age" binding:"required"`
}

// APIPing get api ping
// @Summary Ping
// @Description get api ping
// @Accept  json
// @Produce  json
// @Tags  Ping
// @Security   Beraer
// @Resurce Ping
// @Success 200 {object} main.MessageInfo
// @Router /ping [get]
func APIPing(c *gin.Context) {
	messageinfo := MessageInfo{
		Message: "pong",
	}
	c.JSON(http.StatusOK, messageinfo.RespMessage())
}

// APIHello get api hello
// @Summary Hello
// @Description get api hello
// @Accept  json
// @Produce  json
// @Param   name path string true "name"
// @Tags  Hello
// @Security   Beraer
// @Resurce Hello
// @Success 200 {object} main.MessageInfo
// @Router /hello/{name} [get]
func APIHello(c *gin.Context) {
	messageinfo := MessageInfo{
		Message: c.Param("name"),
	}
	c.JSON(http.StatusOK, messageinfo.RespMessage())
}

// APIParam get api param
// @Summary Param
// @Description get api param
// @Accept json
// @Produce json
// @Param  name query string true "name"
// @Param  age query string true "name"
// @Tags Param
// @Security Beraer
// @Resurce Param
// @Success 200 {object} main.MessageInfo
// @Router /param [get]
func APIParam(c *gin.Context) {
	var param ParamList
	if err := c.ShouldBindQuery(&param); err != nil {
		fmt.Println(err)
		return
	}
	messageinfo := MessageInfo{
		Message: fmt.Sprintf("name: %s , age: %s", param.Name, param.Age),
	}
	c.JSON(http.StatusOK, messageinfo.RespMessage())
}
