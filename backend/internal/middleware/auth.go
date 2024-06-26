package middleware

import (
	"net/http"
	"slices"
	"stp/internal/pkg/config"
	"strings"

	"github.com/gin-gonic/gin"
)

var Auth = &AuthMiddleware{}

type AuthMiddleware struct{}

func (m *AuthMiddleware) Authenticate(ctx *gin.Context) {
	// 某些接口直接放行
	bypassedApis := []string{"ping", "webhooks"}
	if slices.ContainsFunc(bypassedApis, func(str string) bool { return strings.Contains(ctx.Request.URL.String(), str) }) {
		ctx.Next()
		return
	}

	// 获取从 Header 或者 Query 中收到的 Authorization 信息
	authValue := ctx.Request.Header.Get("Authorization")
	if authValue == "" {
		authValue = ctx.Query("Authorization")
	}

	// 校验 Authorization 信息
	if authValue != config.Config.Server.RequestAuth {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "客户端验证不通过",
			"data": nil,
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}
