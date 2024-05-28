package res

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func IsViolatingUniqueConstraint(err error) bool {
	return strings.Contains(strings.ToLower(err.Error()), "unique")
}

// 返回 200 状态码，并携带 data 数据
func Success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
		"data": data,
	})
}

// 返回 400 状态码，参数错误
func ParamErr(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"msg":  "参数错误",
		"data": nil,
	})
}

// 返回 500 状态码，服务器内部错误
func InternalErr(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code": http.StatusInternalServerError,
		"msg":  "服务器内部错误",
		"data": nil,
	})
}
