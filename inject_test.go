package cp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInject(t *testing.T) {
	testData := map[string]string{
		"request-id": "123",
	}

	result := Inject(testData)

	assert.Equal(t, 1, len(result))

	for k, v := range result {
		assert.Equal(t, "baggage-request-id", k)
		assert.Equal(t, "123", v)
	}
}
