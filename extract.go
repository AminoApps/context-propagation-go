package cp

import "strings"

//Extract Convert external context to internal context
func Extract(carrier map[string]string) map[string]string {
	// todo: config propagation by environment variables

	var result map[string]string
	for k, v := range carrier {
		if len(k) > len(BaggagePrefix) && strings.HasPrefix(strings.ToLower(k), BaggagePrefix) {
			if result == nil {
				result = make(map[string]string)
			}
			result[strings.ToLower(k[len(BaggagePrefix):])] = v
		}
	}
	return result
}
