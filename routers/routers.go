package routers

import (
	"go-auth-server/controllers"
	V1Contoller "go-auth-server/controllers/v1"

	"github.com/gin-gonic/gin"
)

func InitializeRouters(r *gin.Engine) {

	// Commmon routers
	health := new(controllers.HealthController)
	r.GET("/health", health.Status)

	// V1 routers
	v1 := r.Group("/v1")
	v1Controller := new(V1Contoller.V1Contoller)

	v1.POST("/signup", v1Controller.Signup)

}
