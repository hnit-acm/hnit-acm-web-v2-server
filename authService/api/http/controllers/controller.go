package controllers

import (
	"auth-service/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hnit-acm/hfunc/hapi"
	"net/http"
)

type AuthServiceController struct {
}

func (s AuthServiceController) RouterRegister(group *gin.RouterGroup) {

}

func (s AuthServiceController) RouterGroupName() (name string) {
	return "auth_service"
}

func (s AuthServiceController) Middlewares() (middlewares []gin.HandlerFunc) {
	return []gin.HandlerFunc{}
}

func (s AuthServiceController) Version() string {
	return "v1"
}

func (s AuthServiceController) Ping() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "ping", "v1", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"msg":  "pong",
			"code": 0,
			"data": "pong",
		})
	}
}

type AuthReq struct {
	Account  string `json:"account,omitempty"`
	Password string `json:"password,omitempty"`
}
type AuthResp struct {
	Token string `json:"token,omitempty"`
}

func (s AuthServiceController) LoginGet() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "login", "v1", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	}
}

func (s AuthServiceController) LoginPost() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodPost, "login", "v1", func(ctx *gin.Context) {
		if true {
			ctx.Redirect(http.StatusFound, "/api/v1/auth_service/auth")
		}
	}
}

func (s AuthServiceController) AuthGet() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodGet, "auth", "v1", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "auth.html", nil)
	}
}

func (s AuthServiceController) Auth() (httpMethod, routeUri, version string, handlerFunc gin.HandlerFunc) {
	return http.MethodPost, "auth", "v1", func(ctx *gin.Context) {
		var (
			err    error
			params = new(AuthReq)
			resp   = new(AuthResp)
		)
		defer func() {
			hapi.JsonResponseErr(ctx, err)
		}()
		if err = ctx.ShouldBind(params); err != nil {
			return
		}
		resp.Token, err = common.JwtEncode(jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: 0,
			Id:        "",
			IssuedAt:  0,
			Issuer:    "",
			NotBefore: 0,
			Subject:   "",
		})
		if err != nil {
			return
		}
		hapi.JsonResponseOk(ctx, resp)
	}
}
