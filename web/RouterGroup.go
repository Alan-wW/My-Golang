package main

import "github.com/gin-gonic/gin"

/*
	将拥有共同URL前缀的路由划分为一个路由组。习惯性一对{}包裹同组的路由
*/
func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {

		})
		userGroup.GET("/login", func(c *gin.Context) {

		})
		userGroup.POST("/login", func(c *gin.Context) {

		})

	}
	/*shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {

		})
		shopGroup.GET("/cart", func(c *gin.Context) {

		})
		shopGroup.POST("/checkout", func(c *gin.Context) {

		})
	}*/

	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {

		})
		shopGroup.GET("/cart", func(c *gin.Context) {

		})
		shopGroup.POST("/checkout", func(c *gin.Context) {

		})
		// 嵌套路由组
		xx := shopGroup.Group("xx")
		xx.GET("/oo", func(c *gin.Context) {

		})
	}
	r.Run()
}
