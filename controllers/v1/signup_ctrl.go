package v1Controller

import "github.com/gin-gonic/gin"

type SignupController struct{}

func (ctrl SignupController) Signup(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "data",
	})
}
