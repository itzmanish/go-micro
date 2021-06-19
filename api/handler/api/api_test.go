package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/itzmanish/go-micro/v2/api"
	"github.com/itzmanish/go-micro/v2/api/handler"
	"github.com/itzmanish/go-micro/v2/api/router/static"
)

func TestAPIHandler(t *testing.T) {
	s := static.NewRouter()
	ep := &api.Endpoint{Name: "some", Path: []string{"/greet"}, Method: []string{"POST"}, Handler: "rpc"}
	s.Register(ep)
	defer s.Deregister(ep)

	apiHandler := NewHandler(handler.WithRouter(s))
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("POST", "/greet", bytes.NewBuffer([]byte(`{"foo":"bar"`)))
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	apiHandler.ServeHTTP(rr, req)
	t.Log(rr.Body.String())

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the response body is what we expect.
	expected := `{"id":"go.micro.client","code":500,"detail":"service some: not found","status":"Internal Server Error"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
