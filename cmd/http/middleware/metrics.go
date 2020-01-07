package middleware

import (
	"os"

	"git.dustess.com/mk-base/gin-ext/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	name := os.Getenv("name")
	if name == "" {
		name = "t_mk_opentrace"
	}

	middleware.Init(name)
}

// PromMiddle prometheus
func PromMiddle() gin.HandlerFunc {
	return middleware.PromMiddleware(nil)
}

// RegisterEndpoint 注册推送接口
func RegisterEndpoint(router *gin.Engine, prefix string) {
	middleware.RegisterEndpoint(router, prefix)
}