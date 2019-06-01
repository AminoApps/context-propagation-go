package context_propagation_go

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	ctx := SetValueToContext(nil, "test-key-useless", "xxx")
	ctx = SetValueToContext(ctx, "test-key", "test-value")
	assert.Equal(t, "test-value", GetValueFromContext(ctx, "test-key"))
}

func TestSetIsolate(t *testing.T) {
	ctx1 := SetValueToContext(nil, "key", "val")
	ctx1 = SetValueToContext(ctx1, "key2", "val1")
	ctx2 := SetValueToContext(ctx1, "key2", "val2")
	ctx3 := SetValueToContext(ctx1, "key2", "val3")
	SetValueToContext(ctx1, "key3", "val4")

	assert.Equal(t, "val", GetValueFromContext(ctx1, "key"))
	assert.Equal(t, "val1", GetValueFromContext(ctx1, "key2"))
	assert.Equal(t, "", GetValueFromContext(ctx1, "key3"))

	assert.Equal(t, "val2", GetValueFromContext(ctx2, "key2"))
	assert.Equal(t, "val3", GetValueFromContext(ctx3, "key2"))
}

func TestSetAndGetNil(t *testing.T) {
	assert.Equal(t, "", GetValueFromContext(context.Background(), "test-key"))
}
