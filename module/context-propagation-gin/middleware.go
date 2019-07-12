package cpgin

import (
	"context"
	"github.com/AminoApps/context-propagation-go"
	"github.com/gin-gonic/gin"
)

//Middleware for gin
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

	carrier := cp.Extract(headersWithFirst)
	if len(carrier) > 0 {
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), cp.InternalContextKey{}, carrier))
	}

	c.Next()
}
