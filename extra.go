package context_propagation_go

import "strings"

func Extract(carrier map[string]string) map[string]string {
	var result map[string]string
	for k, v := range carrier {
		if len(k) > len(BaggagePrefix) && strings.HasPrefix(strings.ToLower(k), BaggagePrefix) {
			if result == nil {
				result = make(map[string]string)
			}
			internalKey := InternalPrefix + strings.ToLower(k[len(BaggagePrefix):])
			result[internalKey] = v
		}
	}
	return result
}
