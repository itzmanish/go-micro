package static

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/itzmanish/go-micro/v2/api"
	"github.com/stretchr/testify/assert"
)

func TestStaticRouter(t *testing.T) {
	ep := &api.Endpoint{Name: "api.greeter.hello", Path: []string{"/greet"}, Method: []string{"POST"}, Handler: "rpc", Host: []string{"", "localhost"}, Stream: true}
	r := NewRouter()
	err := r.Register(ep)
	t.Log(err)
	req, err := http.NewRequest("POST", "/greet", bytes.NewBuffer([]byte(`{"foo":"bar"`)))
	if err != nil {
		t.Fatal(err)
	}
	rep, err := r.Route(req)
	if err != nil {
		t.Fatal(err)
	}
	ep.Name = "greeter.hello"
	assert.Equal(t, ep, rep.Endpoint)
	r.Deregister(ep)
	r.Close()
}
