package main

import (
	"go-auth-server/middlewares"
	"go-auth-server/routers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	r := gin.Default()

	// Setup middlewares
	r.Use(middlewares.RequestIdMiddleware())

	// Init routers
	v1 := r.Group("/v1")

	routers.InitCommonRouter(r)
	routers.InitV1Router(v1)

	r.Run()
}
