package main

import (
	"go-auth-server/middlewares"
	"go-auth-server/routers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	r := gin.Default()

	r.Use(middlewares.RequestIdMiddleware())

	routers.InitializeRouters(r)

	r.Run()
}
