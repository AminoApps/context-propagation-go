package context_propagation_gin

import (
	"context"
	cpg "github.com/AminoApps/context-propagation-go"
	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return handler
}

func handler(c *gin.Context) {
	headersWithFirst := make(map[string]string, len(c.Request.Header))

	for k, v := range c.Request.Header {
		if len(v) > 0 {
			headersWithFirst[k] = v[0]
		}
	}

	carrier := cpg.Extract(headersWithFirst)
	if len(carrier) > 0 {
		c.Set(cpg.InternalContextKey, carrier)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), cpg.InternalContextKey, carrier))
	}

	c.Next()
}
