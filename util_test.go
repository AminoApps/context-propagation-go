package context_propagation_go

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	ctx := SetValueToContext(nil, "test-key", "test-value")
	assert.Equal(t, "test-value", *GetValueFromContext(ctx, "test-key"))
}

func TestSetAndGetNil(t *testing.T) {
	assert.Nil(t, GetValueFromContext(context.Background(), "test-key"))
}
