package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type dummyHandler struct{}

func (*dummyHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestRejectOtherMethods(t *testing.T) {
	var ts = httptest.NewServer(
		http.HandlerFunc(RejectOtherMethods("POST", &dummyHandler{})),
	)
	defer ts.Close()

	rsp, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if rsp.StatusCode != http.StatusMethodNotAllowed {
		t.Error(rsp.Status)
		t.Error("GET method should have been rejected")
	}

	rsp, err = http.Post(ts.URL, "application/json", nil)
	if err != nil {
		t.Error(err)
	}

	if rsp.StatusCode != http.StatusOK {
		t.Error(rsp.Status)
		t.Error("POST method should have been allowed")
	}
}
