package context_propagation_go

import (
	"context"
	"strings"
)

func getInternalKey(key string) string {
	return strings.ToLower(InternalPrefix + key)
}

func GetValueFromContext(c context.Context, key string) *string {
	if c == nil {
		return nil
	}

	val := c.Value(getInternalKey(key))

	if tmp, ok := val.(string); ok {
		return &tmp
	}

	return nil
}

func SetValueToContext(c context.Context, key string, val string) context.Context {
	if c == nil {
		c = context.Background()
	}

	return context.WithValue(c, getInternalKey(key), val)
}
