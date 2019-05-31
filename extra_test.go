package context_propagation_go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtract(t *testing.T) {
	testData := map[string]string{
		"Baggage-Request-Id": "123",
		"other-header":       "456",
	}

	result := Extract(testData)

	assert.Equal(t, 1, len(result))
	for k, v := range result {
		assert.Equal(t, InternalPrefix+"request-id", k)
		assert.Equal(t, "123", v)
	}
}
