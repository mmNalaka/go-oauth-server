package V1Contoller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type V1Contoller struct{}

func (h V1Contoller) Signup(c *gin.Context) {
	data := c.Request.Body

	c.JSON(http.StatusOK, gin.H{"data": data})
}
