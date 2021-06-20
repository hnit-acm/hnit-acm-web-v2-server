package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"user-service/api/controller"
)

func main() {
	hapi.Server("3001", nil, func(c *gin.Engine) {
		hapi.RegisterHandleFunc(c, func(engine *gin.Engine) *gin.RouterGroup {
			return engine.Group("/api")
		}, controller.UserServiceController{})
	})
}
