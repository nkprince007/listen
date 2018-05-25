package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// RejectOtherMethods accepts only the chosen HTTP method
func RejectOtherMethods(method string, h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			// log the requests handled
			defer func(begin time.Time) {
				log.Printf("HTTP path=%s method=%s took=%s",
					r.URL.String(), r.Method, time.Since(begin))
			}(time.Now())

			h.ServeHTTP(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("Only %s requests are allowed", method)))
	}
}
