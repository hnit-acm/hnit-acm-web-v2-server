package main

import (
	controller2 "auth-service/api/http/controller"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
)

func main() {
	hapi.Server("3001", nil, func(c *gin.Engine) {
		hapi.RegisterHandleFunc(c, func(engine *gin.Engine) *gin.RouterGroup {
			return engine.Group("/api")
		}, controller2.AuthServiceController{})
	})
}
