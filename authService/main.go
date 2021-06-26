package main

import (
	"auth-service/api/http/controllers"
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
)

//go:embed static
var static embed.FS

func main() {
	hapi.Server("8010", nil, func(c *gin.Engine) {
		c.LoadHTMLGlob("static/*")
		hapi.RegisterHandleFunc(c, func(engine *gin.Engine) *gin.RouterGroup {
			return engine.Group("/api")
		}, controllers.AuthServiceController{})
	})
}
