package cp

import (
	"context"
)

//GetValueFromContext Get context
func GetValueFromContext(c context.Context, key string) string {
	if c == nil {
		return ""
	}

	carrier := c.Value(InternalContextKey{})

	if tmp, ok := carrier.(map[string]string); ok {
		return tmp[key]
	}
	return ""
}

//SetValueToContext Set context
func SetValueToContext(c context.Context, key string, val string) context.Context {
	if c == nil {
		c = context.Background()
	}

	existCarrier := c.Value(InternalContextKey{})
	if tmp, ok := existCarrier.(map[string]string); ok {
		carrier := make(map[string]string, len(tmp))
		for k, v := range tmp {
			carrier[k] = v
		}
		carrier[key] = val
		return context.WithValue(c, InternalContextKey{}, carrier)
	}

	return context.WithValue(c, InternalContextKey{}, map[string]string{key: val})
}
