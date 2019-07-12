package cp

//InternalContextKey Context key for context propagation
type InternalContextKey struct{}

const (
	//BaggagePrefix Default prefix for opentracing propagation
	BaggagePrefix = "baggage-"
)
