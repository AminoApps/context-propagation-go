package context_propagation_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	e := gin.New()
	e.Use(Middleware())

	requestId := "123456"

	e.GET("/test1", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("%v", c.Request.Context().Value("context-propagation-request-id")))
	})

	e.GET("/test2", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("%v", c.Value("context-propagation-request-id")))
	})

	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://aminoapps.com/test1", nil)
	req.Header.Set("baggage-request-id", requestId)
	req.Header.Set("baggage-other-baggage", "dozer")

	e.ServeHTTP(w1, req)
	e.ServeHTTP(w2, req)

	assert.Equal(t, requestId, w1.Body.String())
	assert.Equal(t, requestId, w2.Body.String())
}
