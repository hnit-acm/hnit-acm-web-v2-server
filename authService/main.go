package main

import (
	"auth-service/api/http/controllers"
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"github.com/hnit-acm/hfunc/hserver/hhttp"
)

//go:embed static
var static embed.FS

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("static/*")
	hapi.RegisterHandleFunc(
		r,
		func(e *gin.Engine) *gin.RouterGroup {
			return e.Group("/api")
		},
		controllers.AuthServiceController{},
	)
	hapi.ServeAny(
		hhttp.WithAddr(":8010"),
		hhttp.WithHandler(r),
	)
}
