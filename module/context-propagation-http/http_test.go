package cphttp

import (
	"github.com/AminoApps/context-propagation-go"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context/ctxhttp"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/test", Wrap(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(cp.GetValueFromContext(req.Context(), "request-id")))
	})))
	server := httptest.NewServer(mux)
	defer server.Close()

	client := WrapClient(&http.Client{})
	ctx := cp.SetValueToContext(nil, "request-id", "123")

	resp, err := ctxhttp.Get(ctx, client, server.URL+"/test")
	assert.Nil(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, "123", string(body))
}
