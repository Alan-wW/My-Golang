package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		//获取参数
		username := c.Query("username")
		//获取参数，如果没有则取默认值
		hobby := c.DefaultQuery("hobby", "篮球")
		c.JSON(200, gin.H{
			"message":  "你好",
			"username": username,
			"hobby":    hobby,
		})
	})
	r.Run()
}
