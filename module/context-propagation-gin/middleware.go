package context_propagation_gin

import (
	cpg "github.com/AminoApps/context-propagation-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
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

	ctx := c.Request.Context()

	for k, v := range cpg.Extract(headersWithFirst) {
		c.Set(k, v)
		ctx = context.WithValue(ctx, k, v)
	}

	c.Request = c.Request.WithContext(ctx)

	c.Next()
}
