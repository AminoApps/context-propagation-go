package context_propagation_http

import (
	cpg "github.com/AminoApps/context-propagation-go"
	"net/http"
)

func WrapClient(c *http.Client) *http.Client {
	if c == nil {
		c = http.DefaultClient
	}
	copied := *c

	if copied.Transport == nil {
		copied.Transport = &roundTripper{rt: http.DefaultTransport}
	} else {
		copied.Transport = &roundTripper{rt: copied.Transport}
	}
	return &copied
}

type roundTripper struct {
	rt http.RoundTripper
}

func (s *roundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	carrier := r.Context().Value(cpg.InternalContextKey)
	headers := cpg.Inject(carrier)

	for k, v := range headers {
		r.Header.Set(k, v)
	}

	return s.rt.RoundTrip(r)
}
