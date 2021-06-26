package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func AuthMW(ctx *gin.Context) {
	accessToken, err := ctx.Cookie("access_token")
	if err != nil || accessToken == "" {
		//hapi.JsonResponseErr(ctx,err,hapi.WithStatus(http.StatusUnauthorized))

		redirectParam := ctx.Request.URL
		redirectParam.Host = ctx.Request.Host
		redirectUrl := url.URL{
			Scheme:      "http",
			Opaque:      "",
			User:        nil,
			Host:        "127.0.0.1:8010",
			Path:        "/api/v1/auth_service/login",
			RawPath:     "",
			ForceQuery:  false,
			RawQuery:    "redirect=" + redirectParam.String(),
			Fragment:    "",
			RawFragment: "",
		}

		ctx.Redirect(http.StatusFound, redirectUrl.String())
		ctx.Abort()
		return
	}
	ctx.Next()
}
