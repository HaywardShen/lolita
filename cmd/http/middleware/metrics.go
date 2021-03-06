package middleware

import (
	"os"

	metrics "github.com/asppj/lolita/ext/gin-metrics"

	"github.com/gin-gonic/gin"
)

func init() {
	name := os.Getenv("name")
	if name == "" {
		name = "lolita"
	}
	metrics.Init(name)
	// middleware.Init(name)
}

// PromMiddle prometheus
func PromMiddle() gin.HandlerFunc {
	// return middleware.PromMiddleware(nil)
	return metrics.PromMiddleware(nil)
}

// RegisterEndpoint 注册推送接口
func RegisterEndpoint(router *gin.Engine, prefix string) {
	metrics.RegisterEndpoint(router, prefix)
}
