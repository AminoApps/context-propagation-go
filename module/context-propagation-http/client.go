package cphttp

import (
	"github.com/AminoApps/context-propagation-go"
	"net/http"
)

//Wrap http.Client
func WrapClient(c *http.Client) *http.Client {
	if c == nil {
		c = http.DefaultClient
	}
	copied := *c

	copied.Transport = WrapRoundTripper(copied.Transport)

	return &copied
}

//Wrap http.RoundTripper
func WrapRoundTripper(r http.RoundTripper) http.RoundTripper {
	if r == nil {
		r = http.DefaultTransport
	}
	return &roundTripper{rt: r}
}

type roundTripper struct {
	rt http.RoundTripper
}

func (s *roundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	carrier := r.Context().Value(cp.InternalContextKey{})
	headers := cp.Inject(carrier)

	for k, v := range headers {
		r.Header.Set(k, v)
	}

	return s.rt.RoundTrip(r)
}
