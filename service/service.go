package service

import (
	//"github.com/go-martini/martini"
	"github.com/gin-gonic/gin"
)

func NewServer(port string) { /*新建服务器*/
	r := gin.Default()
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello "+name)
	})
	r.Run(":" + port)
}
