package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiv1 := r.Group("api/v1")

	apiv1.POST("/signup", func(c *gin.Context) {
		data := c.Request.Body
		print(data)

		c.JSON(200, gin.H{
			"_id":            "58457fe6b27...",
			"email_verified": false,
			"email":          "test.account@signup.com",
			"username":       "johndoe",
			"given_name":     "John",
			"family_name":    "Doe",
			"name":           "John Doe",
			"nickname":       "johnny",
			"picture":        "http://example.org/jdoe.png",
		})
	})

	r.Run()
}
