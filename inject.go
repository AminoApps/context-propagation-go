package context_propagation_go

func Inject(carrier interface{}) map[string]string {
	// todo: config propagation by environment variables

	if tmp, ok := carrier.(map[string]string); ok {
		var result map[string]string
		for k, v := range tmp {
			if result == nil {
				result = make(map[string]string)
			}
			result[BaggagePrefix+k] = v
		}
		return result
	}

	return nil
}
