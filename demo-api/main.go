package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/")

	r.Run
}

func welcome() func(c *gin.COntext) {
	return func(c *gin.COntext) {
		c.JSON(200, gin.H{
			"Message" : "Hello",
		})
	}
}
