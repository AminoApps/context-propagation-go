package cphttp

import (
	"context"
	"github.com/AminoApps/context-propagation-go"
	"net/http"
)

//Wrap http.Handler
func Wrap(h http.Handler) http.Handler {
	if h == nil {
		panic("h == nil")
	}

	return &handler{
		handler: h,
	}
}

type handler struct {
	handler http.Handler
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	headersWithFirst := make(map[string]string, len(req.Header))

	for k, v := range req.Header {
		if len(v) > 0 {
			headersWithFirst[k] = v[0]
		}
	}

	carrier := cp.Extract(headersWithFirst)
	if len(carrier) > 0 {
		req = req.WithContext(context.WithValue(req.Context(), cp.InternalContextKey{}, carrier))
	}

	h.handler.ServeHTTP(w, req)
}
