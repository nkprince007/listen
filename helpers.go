package main

import (
	"fmt"
	"net/http"
)

// RejectOtherMethods accepts only the chosen HTTP method
func RejectOtherMethods(method string, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			h.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("Only %s requests are allowed", method)))
	}
}
