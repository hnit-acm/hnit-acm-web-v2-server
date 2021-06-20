package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserServiceController struct {
}

func (s UserServiceController) RouterRegister(group *gin.RouterGroup) {

}

func (s UserServiceController) RouterGroupName() (name string) {
	return "user-service"
}

func (s UserServiceController) Middlewares() (middlewares []gin.HandlerFunc) {
	return []gin.HandlerFunc{}
}

func (s UserServiceController) Version() string {
	return "v1"
}

func (s UserServiceController) Ping() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "ping", "v1", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"msg":  "pong",
			"code": 0,
			"data": "pong",
		})
	}
}
