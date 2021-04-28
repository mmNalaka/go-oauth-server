package routers

import (
	"go-auth-server/controllers"

	"github.com/gin-gonic/gin"
)

func InitCommonRouter(r *gin.Engine) {

	health := new(controllers.HealthController)
	r.GET("/health", health.Status)
}
