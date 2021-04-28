package routers

import (
	v1Controller "go-auth-server/controllers/v1"

	"github.com/gin-gonic/gin"
)

func InitV1Router(v1 *gin.RouterGroup) {
	v1Controller := new(v1Controller.SignupController)

	v1.GET("/signup", v1Controller.Signup)
}
